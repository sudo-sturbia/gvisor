// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package mq provides an implementation for POSIX message queues.
package mq

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/context"
	"gvisor.dev/gvisor/pkg/errors/linuxerr"
	"gvisor.dev/gvisor/pkg/marshal/primitive"
	"gvisor.dev/gvisor/pkg/sentry/fs"
	"gvisor.dev/gvisor/pkg/sentry/kernel/auth"
	"gvisor.dev/gvisor/pkg/sentry/socket/netlink"
	"gvisor.dev/gvisor/pkg/sentry/vfs"
	"gvisor.dev/gvisor/pkg/sync"
	"gvisor.dev/gvisor/pkg/waiter"
)

// AccessType is the access type passed to mq_open.
type AccessType int

// Possible access types.
const (
	ReadOnly AccessType = iota
	WriteOnly
	ReadWrite
)

const (
	MaxName     = 255                   // Maximum size for a queue name.
	maxPriority = linux.MQ_PRIO_MAX - 1 // Highest possible message priority.

	maxQueuesDefault = linux.DFLT_QUEUESMAX // Default max number of queues.

	maxMsgDefault   = linux.DFLT_MSG    // Default max number of messages per queue.
	maxMsgMin       = linux.MIN_MSGMAX  // Min value for max number of messages per queue.
	maxMsgLimit     = linux.DFLT_MSGMAX // Limit for max number of messages per queue.
	maxMsgHardLimit = linux.HARD_MSGMAX // Hard limit for max number of messages per queue.

	msgSizeDefault   = linux.DFLT_MSGSIZE    // Default max message size.
	msgSizeMin       = linux.MIN_MSGSIZEMAX  // Min value for max message size.
	msgSizeLimit     = linux.DFLT_MSGSIZEMAX // Limit for max message size.
	msgSizeHardLimit = linux.HARD_MSGSIZEMAX // Hard limit for max message size.
)

// Registry is a POSIX message queue registry.
//
// Unlike SysV utilities, Registry is not map-based. It uses a provided
// RegistryImpl backed by a virtual filesystem to implement registry operations.
//
// +stateify savable
type Registry struct {
	// userNS is the user namespace containing this registry. Immutable.
	userNS *auth.UserNamespace

	// mu protects all fields below.
	mu sync.Mutex `state:"nosave"`

	// impl is an implementation of several message queue utilities needed by
	// the registry. impl should be provided by mqfs.
	impl RegistryImpl
}

// RegistryImpl defines utilities needed by a Registry to provide actual
// registry implementation. It works mainly as an abstraction layer used by
// Registry to avoid dealing directly with the filesystem. RegistryImpl should
// be implemented by mqfs and provided to Registry at initialization.
type RegistryImpl interface {
	// Get searchs for a queue with the given name, if it exists, the queue is
	// used to create a new FD, return it and return true. If the queue  doesn't
	// exist, return false and no error. An error is returned if creation fails.
	Get(ctx context.Context, name string, access AccessType, block bool, flags uint32) (*vfs.FileDescription, bool, error)

	// New creates a new inode and file description using the given queue,
	// inserts the inode into the filesystem tree using the given name, and
	// returns the file description. An error is returned if creation fails, or
	// if the name already exists.
	New(ctx context.Context, name string, q *Queue, access AccessType, block bool, perm linux.FileMode, flags uint32) (*vfs.FileDescription, error)

	// Unlink removes the queue with given name from the registry, and returns
	// an error if the name doesn't exist.
	Unlink(ctx context.Context, name string) error

	// Destroy destroys the registry.
	Destroy(context.Context)
}

// NewRegistry returns a new, initialized message queue registry. NewRegistry
// should be called when a new message queue filesystem is created, once per
// IPCNamespace.
func NewRegistry(userNS *auth.UserNamespace, impl RegistryImpl) *Registry {
	return &Registry{
		userNS: userNS,
		impl:   impl,
	}
}

// OpenOpts holds the options passed to FindOrCreate.
type OpenOpts struct {
	Name      string
	Access    AccessType
	Create    bool
	Exclusive bool
	Block     bool
}

// FindOrCreate creates a new POSIX message queue or opens an existing queue.
// See mq_open(2).
func (r *Registry) FindOrCreate(ctx context.Context, opts OpenOpts, perm linux.FileMode, attr *linux.MqAttr) (*vfs.FileDescription, error) {
	// mq_overview(7) mentions that: "Each message queue is identified by a name
	// of the form '/somename'", but the mq_open(3) man pages mention:
	//   "The mq_open() library function is implemented on top of a system call
	//    of the same name.  The library function performs the check that the
	//    name starts with a slash (/), giving the EINVAL error if it does not.
	//    The kernel system call expects name to contain no preceding slash, so
	//    the C library function passes name without the preceding slash (i.e.,
	//    name+1) to the system call."
	// So we don't need to check it.

	if len(opts.Name) == 0 {
		return nil, linuxerr.ENOENT
	}
	if len(opts.Name) > MaxName {
		return nil, linuxerr.ENAMETOOLONG
	}
	if strings.ContainsRune(opts.Name, '/') {
		return nil, linuxerr.EACCES
	}
	if opts.Name == "." || opts.Name == ".." {
		return nil, linuxerr.EINVAL
	}

	// Construct status flags.
	var flags uint32
	if opts.Block {
		flags = linux.O_NONBLOCK
	}
	switch opts.Access {
	case ReadOnly:
		flags = flags | linux.O_RDONLY
	case WriteOnly:
		flags = flags | linux.O_WRONLY
	case ReadWrite:
		flags = flags | linux.O_RDWR
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	fd, ok, err := r.impl.Get(ctx, opts.Name, opts.Access, opts.Block, flags)
	if err != nil {
		return nil, err
	}

	if ok {
		if opts.Create && opts.Exclusive {
			// "Both O_CREAT and O_EXCL were specified in oflag, but a queue
			//  with this name already exists."
			return nil, linuxerr.EEXIST
		}
		return fd, nil
	}

	if !opts.Create {
		// "The O_CREAT flag was not specified in oflag, and no queue with this name
		//  exists."
		return nil, linuxerr.ENOENT
	}

	q, err := r.newQueueLocked(auth.CredentialsFromContext(ctx), fs.FileOwnerFromContext(ctx), fs.FilePermsFromMode(perm), attr)
	if err != nil {
		return nil, err
	}
	return r.impl.New(ctx, opts.Name, q, opts.Access, opts.Block, perm, flags)
}

// newQueueLocked creates a new queue using the given attributes. If attr is nil
// return a queue with default values, otherwise use attr to create a new queue,
// and return an error if attributes are invalid.
func (r *Registry) newQueueLocked(creds *auth.Credentials, owner fs.FileOwner, perms fs.FilePermissions, attr *linux.MqAttr) (*Queue, error) {
	if attr == nil {
		return &Queue{
			owner:           owner,
			perms:           perms,
			maxMessageCount: int64(maxMsgDefault),
			maxMessageSize:  uint64(msgSizeDefault),
		}, nil
	}

	// "O_CREAT was specified in oflag, and attr was not NULL, but
	//  attr->mq_maxmsg or attr->mq_msqsize was invalid.  Both of these fields
	//  these fields must be greater than zero.  In a process that is
	//  unprivileged (does not have the CAP_SYS_RESOURCE capability),
	//  attr->mq_maxmsg must be less than or equal to the msg_max limit, and
	//  attr->mq_msgsize must be less than or equal to the msgsize_max limit.
	//  In addition, even in a privileged process, attr->mq_maxmsg cannot
	//  exceed the HARD_MAX limit." - man mq_open(3).
	if attr.MqMaxmsg <= 0 || attr.MqMsgsize <= 0 {
		return nil, linuxerr.EINVAL
	}

	if attr.MqMaxmsg > maxMsgHardLimit || (!creds.HasCapabilityIn(linux.CAP_SYS_RESOURCE, r.userNS) && (attr.MqMaxmsg > maxMsgLimit || attr.MqMsgsize > msgSizeLimit)) {
		return nil, linuxerr.EINVAL
	}

	return &Queue{
		owner:           owner,
		perms:           perms,
		maxMessageCount: attr.MqMaxmsg,
		maxMessageSize:  uint64(attr.MqMsgsize),
	}, nil
}

// Remove removes the queue with the given name from the registry. See
// mq_unlink(2).
func (r *Registry) Remove(ctx context.Context, name string) error {
	if len(name) > MaxName {
		return linuxerr.ENAMETOOLONG
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	return r.impl.Unlink(ctx, name)
}

// Destroy destroys the registry and releases all held references.
func (r *Registry) Destroy(ctx context.Context) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.impl.Destroy(ctx)
}

// Impl returns RegistryImpl inside r.
func (r *Registry) Impl() RegistryImpl {
	return r.impl
}

// Queue represents a POSIX message queue.
//
// +stateify savable
type Queue struct {
	// owner is the registry's owner. Immutable.
	owner fs.FileOwner

	// perms is the registry's access permissions. Immutable.
	perms fs.FilePermissions

	// mu protects all the fields below.
	mu sync.Mutex `state:"nosave"`

	// senders is a queue of currently blocked senders. Senders are notified
	// when space isi available in the queue for a new message.
	senders waiter.Queue

	// receivers is a queue of currently blocked receivers. Receivers are
	// notified when a new message is inserted in the queue.
	receivers waiter.Queue

	// messages is a list of messages currently in the queue.
	messages msgList

	// subscriber represents a task registered to receive async notification
	// from this queue.
	subscriber *Subscriber

	// messageCount is the number of messages currently in the queue.
	messageCount int64

	// maxMessageCount is the maximum number of messages that the queue can
	// hold.
	maxMessageCount int64

	// maxMessageSize is the maximum size of a message held by the queue.
	maxMessageSize uint64

	// byteCount is the number of bytes of data in all messages in the queue.
	byteCount uint64
}

// Message holds a message exchanged through a Queue via mq_timedsend(2) and
// mq_timedreceive(2), and additional info relating to the message.
//
// +stateify savable
type Message struct {
	msgEntry

	// Text is the message's sent content.
	Text string

	// Size is the message's size in bytes.
	Size uint64

	// Priority is the message's priority.
	Priority uint32
}

// Target is used by Subscriber to send a signal to a task when a message
// arrives. kernel.Task is not directly used to prevent circular dependency.
type Target interface {
	SendSignal(info *linux.SignalInfo) error
}

// Subscriber represents a task registered for async notification from a Queue.
//
// +stateify savable
type Subscriber struct {
	// target represents a task registered to receive a signal.
	target Target

	// signal is the signal number to be sent.
	signal linux.Signal

	// cookie is a 32-byte cookie, sent as a trigger to a netlink socket when
	// mq_notify(SIGEV_THREAD) is used.
	cookie []byte

	// method is the method of notification.
	method int32

	// pid is the PID of the registered task.
	pid int32
}

// Blocker is used for blocking Queue.Send, and Queue.Receive calls, and serves
// as an abstracted version of kernel.Task. kernel.Task is not directly used to
// prevent circular dependency.
type Blocker interface {
	BlockWithTimeout(C chan struct{}, haveTimeout bool, timeout time.Duration) (time.Duration, error)
}

// send adds a given message to the queue, and returns an error if sending
// fails. See mq_timedsend(2).
func (q *Queue) send(ctx context.Context, msg Message, b Blocker, timeout time.Duration, block bool) error {
	// Fast path: attempt a non-blocking push.
	if err := q.push(ctx, msg); err != linuxerr.EWOULDBLOCK {
		return err
	}

	if !block {
		return linuxerr.EAGAIN
	}
	if timeout == 0 {
		return linuxerr.ETIMEDOUT
	}

	// Slow path: the queue was found to be full, and we were asked to block.

	e, ch := waiter.NewChannelEntry(nil)
	q.senders.EventRegister(&e, waiter.EventOut)
	defer q.senders.EventUnregister(&e)

	// We need to check again before blocking since space may have become
	// available.
	for {
		if err := q.push(ctx, msg); err != linuxerr.EWOULDBLOCK {
			return err
		}

		remaining, err := b.BlockWithTimeout(ch, timeout >= 0, timeout)
		if err != nil {
			return err
		}
		timeout = remaining
	}
}

// push adds a message to the queue's message list based on its priority, and
// notifies waiting receivers that a message has been inserted. An error is
// returned if adding the message would cause the queue to exceed max capacity.
func (q *Queue) push(ctx context.Context, msg Message) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Unlike SysV, POSIX message queues only care about the overall message
	// count, but not the overall size.
	if q.messageCount >= q.maxMessageCount {
		return linuxerr.EWOULDBLOCK
	}

	// mq_send(3) man pages mention:
	//   "Messages are placed on the queue in decreasing order of priority,
	//    with newer messages of the same priority being placed after older
	//    messages with the same priority."
	// To make retreival easier, we arrange messages on insertion.

	// Insert at the back if the queue is empty, or if all messages have a
	// higher priority.
	if q.messages.Empty() || q.messages.Back().Priority >= msg.Priority {
		q.messages.PushBack(&msg)
	} else {
		entry := q.messages.Front()
		for entry.Priority >= msg.Priority {
			entry = entry.Next()
		}

		// We already checked if we should be at the back, so we don't have to
		// worry about entry being nil.
		q.messages.InsertBefore(entry, &msg)
	}

	q.messageCount++
	q.byteCount += msg.Size

	// Notify subscriber if this is the first message and there are no
	// waiting receivers.
	if q.subscriber != nil && q.receivers.IsEmpty() && q.messageCount == 1 {
		switch q.subscriber.method {
		case linux.SIGEV_NONE:
			// "A "null" notification: the calling process is registered as the
			//  target for notification, but when a message arrives, no
			//  notification is sent."
		case linux.SIGEV_SIGNAL:
			q.notifySignal(ctx)
		case linux.SIGEV_THREAD:
		}
		q.subscriber = nil
	}

	// Notify waiting receivers.
	q.receivers.Notify(waiter.EventIn)

	return nil
}

// notifySignal sends an async signal to a task, see mq_notify(SIGEV_SIGNAL).
func (q *Queue) notifySignal(ctx context.Context) {
	pid, ok := context.ThreadGroupIDFromContext(ctx)
	if !ok {
		panic("mq.push: failed to get pid")
	}

	info := &linux.SignalInfo{
		Signo: int32(q.subscriber.signal),
		Code:  linux.SI_MESGQ,
	}
	info.SetPID(pid)
	info.SetUID(int32(fs.FileOwnerFromContext(ctx).UID))
	q.subscriber.target.SendSignal(info)
}

// FDTable is used by Queue to get a netlink socket with a given FD. kernel.Task
// can't be directly used to prevent circular dependency.
type FDTable interface {
	GetFileVFS2(fd int32) *vfs.FileDescription
}

// notifyThread implements the behaviour expected by mq_notify(SIGEV_THREAD).
func (q *Queue) notifyThread(ctx context.Context, t FDTable, notifyCode uint) {
	// Get socket from the file descriptor.
	file := t.GetFileVFS2(int32(q.subscriber.signal))
	if file == nil {
		panic("mq.notifyThread: sigev_signo is not a file descriptor")
	}
	defer file.DecRef(ctx)

	// Extract the socket.
	s, ok := file.Impl().(*netlink.SocketVFS2)
	if !ok {
		panic("mq.notifyThread: file descriptor is not a socket")
	}

	// Create message to send.
	q.subscriber.cookie[len(q.subscriber.cookie)-1] = byte(notifyCode)

	var message netlink.Message
	message.Put(primitive.ByteSlice(q.subscriber.cookie))

	set := &netlink.MessageSet{
		Messages: []*netlink.Message{&message},
	}
	s.SendResponse(ctx, set)
}

// receive returns the message with the highest priority from the queue. See
// mq_timedreceive(2).
func (q *Queue) receive(ctx context.Context, b Blocker, timeout time.Duration, block bool) (*Message, error) {
	// Fast path: attempt a non-blocking pop.
	if msg, err := q.pop(ctx); err != linuxerr.EWOULDBLOCK {
		return msg, err
	}

	if !block {
		return nil, linuxerr.EAGAIN
	}
	if timeout == 0 {
		return nil, linuxerr.ETIMEDOUT
	}

	// Slow path: the queue was found to be empty, and we were asked to block.

	e, ch := waiter.NewChannelEntry(nil)
	q.receivers.EventRegister(&e, waiter.EventIn)
	defer q.receivers.EventUnregister(&e)

	// We need to check again before blocking since space may have become
	// available.
	for {
		if msg, err := q.pop(ctx); err != linuxerr.EWOULDBLOCK {
			return msg, err
		}

		remaining, err := b.BlockWithTimeout(ch, timeout >= 0, timeout)
		if err != nil {
			return nil, err
		}
		timeout = remaining
	}
}

// pop removes the highest priority message from the queue and returns it. An
// error EWOULDBLOCK is returned if the queue is empty.
func (q *Queue) pop(ctx context.Context) (*Message, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.messages.Empty() {
		return nil, linuxerr.EWOULDBLOCK
	}

	// Messages are ordered on insertion, so the head of the linked list always
	// has the highest priority.
	msg := q.messages.Front()

	q.messages.Remove(msg)
	q.messageCount--
	q.byteCount -= msg.Size

	// Notify waiting senders.
	q.senders.Notify(waiter.EventOut)

	return msg, nil
}

// Attr returns queue's attributes, including view's O_NONBLOCK flag. See
// mq_getattr(3).
func (q *Queue) Attr(block bool) *linux.MqAttr {
	q.mu.Lock()
	defer q.mu.Unlock()

	var nonBlock int64
	if !block {
		nonBlock = linux.O_NONBLOCK
	}

	return &linux.MqAttr{
		MqFlags:   nonBlock,
		MqMaxmsg:  int64(q.maxMessageCount),
		MqMsgsize: int64(q.maxMessageSize),
		MqCurmsgs: int64(q.messageCount),
	}
}

// Register implements View.Register.
func (q *Queue) Register(t Target, event *linux.Sigevent, cookie []byte, pid int32) error {
	if invalidEvent(event) {
		// "sevp->sigev_notify is not one of the permitted values; or
		//  sevp->sigev_notify is SIGEV_SIGNAL and sevp->sigev_signo is not a
		//  valid signal number."
		return linuxerr.EINVAL
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	if q.subscriber != nil {
		// "Another process has already registered to receive notification for
		//  this message queue."
		return linuxerr.EBUSY
	}

	subscriber := &Subscriber{
		signal: linux.Signal(event.Signo),
		method: event.Notify,
		pid:    pid,
	}

	switch event.Notify {
	case linux.SIGEV_SIGNAL:
		subscriber.target = t
	case linux.SIGEV_THREAD:
		subscriber.cookie = cookie
	}

	q.subscriber = subscriber
	return nil
}

// Unregister implements View.Unregister.
func (q *Queue) Unregister(pid int32) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.subscriber != nil && pid == q.subscriber.pid {
		q.subscriber = nil
	}
}

// invalidEvent return true if given event is invalid.
func invalidEvent(event *linux.Sigevent) bool {
	return event.Notify != linux.SIGEV_NONE && (event.Notify != linux.SIGEV_SIGNAL || (event.Notify == linux.SIGEV_SIGNAL && !linux.Signal(event.Signo).IsValid())) && event.Notify != linux.SIGEV_THREAD
}

// Generate implements vfs.DynamicBytesSource.Generate. Queue is used as a
// DynamicBytesSource for mqfs's QueueInode.
func (q *Queue) Generate(ctx context.Context, buf *bytes.Buffer) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	var (
		pid    int32
		method int32
		signal linux.Signal
	)
	if q.subscriber != nil {
		pid = q.subscriber.pid
		method = q.subscriber.method
		signal = q.subscriber.signal
	}

	buf.WriteString(
		fmt.Sprintf("QSIZE:%-10d NOTIFY:%-5d SIGNO:%-5d NOTIFY_PID:%-6d\n",
			q.byteCount, method, signal, pid),
	)
	return nil
}

// Flush implements View.Flush.
func (q *Queue) Flush(ctx context.Context) {
	q.mu.Lock()
	defer q.mu.Unlock()

	pid, ok := context.ThreadGroupIDFromContext(ctx)
	if ok {
		if q.subscriber != nil && pid == q.subscriber.pid {
			q.subscriber = nil
		}
	}
}

// Readiness implements Waitable.Readiness.
func (q *Queue) Readiness(mask waiter.EventMask) waiter.EventMask {
	q.mu.Lock()
	defer q.mu.Unlock()

	events := waiter.EventMask(0)
	if q.messageCount > 0 {
		events |= waiter.ReadableEvents
	}
	if q.messageCount < q.maxMessageCount {
		events |= waiter.WritableEvents
	}
	return events & mask
}

// EventRegister implements Waitable.EventRegister.
func (q *Queue) EventRegister(e *waiter.Entry, mask waiter.EventMask) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if mask&waiter.WritableEvents != 0 {
		q.senders.EventRegister(e, waiter.EventOut)
	}
	if mask&waiter.ReadableEvents != 0 {
		q.receivers.EventRegister(e, waiter.EventIn)
	}
}

// EventUnregister implements Waitable.EventUnregister.
func (q *Queue) EventUnregister(e *waiter.Entry) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.senders.EventUnregister(e)
	q.receivers.EventUnregister(e)
}

// HasPermissions returns true if the given credentials meet the access
// permissions required by the queue.
func (q *Queue) HasPermissions(creds *auth.Credentials, req fs.PermMask) bool {
	p := q.perms.Other
	if q.owner.UID == creds.EffectiveKUID {
		p = q.perms.User
	} else if creds.InGroup(q.owner.GID) {
		p = q.perms.Group
	}
	return p.SupersetOf(req)
}

// MaxMsgSize implements View.MaxMsgSize.
func (q *Queue) MaxMsgSize() uint64 {
	return q.maxMessageSize
}

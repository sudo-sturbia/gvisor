// automatically generated by stateify.

package fs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *StableAttr) StateTypeName() string {
	return "pkg/sentry/fs.StableAttr"
}

func (s *StableAttr) StateFields() []string {
	return []string{
		"Type",
		"DeviceID",
		"InodeID",
		"BlockSize",
		"DeviceFileMajor",
		"DeviceFileMinor",
	}
}

func (s *StableAttr) beforeSave() {}

func (s *StableAttr) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Type)
	stateSinkObject.Save(1, &s.DeviceID)
	stateSinkObject.Save(2, &s.InodeID)
	stateSinkObject.Save(3, &s.BlockSize)
	stateSinkObject.Save(4, &s.DeviceFileMajor)
	stateSinkObject.Save(5, &s.DeviceFileMinor)
}

func (s *StableAttr) afterLoad() {}

func (s *StableAttr) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Type)
	stateSourceObject.Load(1, &s.DeviceID)
	stateSourceObject.Load(2, &s.InodeID)
	stateSourceObject.Load(3, &s.BlockSize)
	stateSourceObject.Load(4, &s.DeviceFileMajor)
	stateSourceObject.Load(5, &s.DeviceFileMinor)
}

func (u *UnstableAttr) StateTypeName() string {
	return "pkg/sentry/fs.UnstableAttr"
}

func (u *UnstableAttr) StateFields() []string {
	return []string{
		"Size",
		"Usage",
		"Perms",
		"Owner",
		"AccessTime",
		"ModificationTime",
		"StatusChangeTime",
		"Links",
	}
}

func (u *UnstableAttr) beforeSave() {}

func (u *UnstableAttr) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.Size)
	stateSinkObject.Save(1, &u.Usage)
	stateSinkObject.Save(2, &u.Perms)
	stateSinkObject.Save(3, &u.Owner)
	stateSinkObject.Save(4, &u.AccessTime)
	stateSinkObject.Save(5, &u.ModificationTime)
	stateSinkObject.Save(6, &u.StatusChangeTime)
	stateSinkObject.Save(7, &u.Links)
}

func (u *UnstableAttr) afterLoad() {}

func (u *UnstableAttr) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.Size)
	stateSourceObject.Load(1, &u.Usage)
	stateSourceObject.Load(2, &u.Perms)
	stateSourceObject.Load(3, &u.Owner)
	stateSourceObject.Load(4, &u.AccessTime)
	stateSourceObject.Load(5, &u.ModificationTime)
	stateSourceObject.Load(6, &u.StatusChangeTime)
	stateSourceObject.Load(7, &u.Links)
}

func (a *AttrMask) StateTypeName() string {
	return "pkg/sentry/fs.AttrMask"
}

func (a *AttrMask) StateFields() []string {
	return []string{
		"Type",
		"DeviceID",
		"InodeID",
		"BlockSize",
		"Size",
		"Usage",
		"Perms",
		"UID",
		"GID",
		"AccessTime",
		"ModificationTime",
		"StatusChangeTime",
		"Links",
	}
}

func (a *AttrMask) beforeSave() {}

func (a *AttrMask) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.Type)
	stateSinkObject.Save(1, &a.DeviceID)
	stateSinkObject.Save(2, &a.InodeID)
	stateSinkObject.Save(3, &a.BlockSize)
	stateSinkObject.Save(4, &a.Size)
	stateSinkObject.Save(5, &a.Usage)
	stateSinkObject.Save(6, &a.Perms)
	stateSinkObject.Save(7, &a.UID)
	stateSinkObject.Save(8, &a.GID)
	stateSinkObject.Save(9, &a.AccessTime)
	stateSinkObject.Save(10, &a.ModificationTime)
	stateSinkObject.Save(11, &a.StatusChangeTime)
	stateSinkObject.Save(12, &a.Links)
}

func (a *AttrMask) afterLoad() {}

func (a *AttrMask) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.Type)
	stateSourceObject.Load(1, &a.DeviceID)
	stateSourceObject.Load(2, &a.InodeID)
	stateSourceObject.Load(3, &a.BlockSize)
	stateSourceObject.Load(4, &a.Size)
	stateSourceObject.Load(5, &a.Usage)
	stateSourceObject.Load(6, &a.Perms)
	stateSourceObject.Load(7, &a.UID)
	stateSourceObject.Load(8, &a.GID)
	stateSourceObject.Load(9, &a.AccessTime)
	stateSourceObject.Load(10, &a.ModificationTime)
	stateSourceObject.Load(11, &a.StatusChangeTime)
	stateSourceObject.Load(12, &a.Links)
}

func (p *PermMask) StateTypeName() string {
	return "pkg/sentry/fs.PermMask"
}

func (p *PermMask) StateFields() []string {
	return []string{
		"Read",
		"Write",
		"Execute",
	}
}

func (p *PermMask) beforeSave() {}

func (p *PermMask) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Read)
	stateSinkObject.Save(1, &p.Write)
	stateSinkObject.Save(2, &p.Execute)
}

func (p *PermMask) afterLoad() {}

func (p *PermMask) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Read)
	stateSourceObject.Load(1, &p.Write)
	stateSourceObject.Load(2, &p.Execute)
}

func (f *FilePermissions) StateTypeName() string {
	return "pkg/sentry/fs.FilePermissions"
}

func (f *FilePermissions) StateFields() []string {
	return []string{
		"User",
		"Group",
		"Other",
		"Sticky",
		"SetUID",
		"SetGID",
	}
}

func (f *FilePermissions) beforeSave() {}

func (f *FilePermissions) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.User)
	stateSinkObject.Save(1, &f.Group)
	stateSinkObject.Save(2, &f.Other)
	stateSinkObject.Save(3, &f.Sticky)
	stateSinkObject.Save(4, &f.SetUID)
	stateSinkObject.Save(5, &f.SetGID)
}

func (f *FilePermissions) afterLoad() {}

func (f *FilePermissions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.User)
	stateSourceObject.Load(1, &f.Group)
	stateSourceObject.Load(2, &f.Other)
	stateSourceObject.Load(3, &f.Sticky)
	stateSourceObject.Load(4, &f.SetUID)
	stateSourceObject.Load(5, &f.SetGID)
}

func (f *FileOwner) StateTypeName() string {
	return "pkg/sentry/fs.FileOwner"
}

func (f *FileOwner) StateFields() []string {
	return []string{
		"UID",
		"GID",
	}
}

func (f *FileOwner) beforeSave() {}

func (f *FileOwner) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.UID)
	stateSinkObject.Save(1, &f.GID)
}

func (f *FileOwner) afterLoad() {}

func (f *FileOwner) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.UID)
	stateSourceObject.Load(1, &f.GID)
}

func (d *DentAttr) StateTypeName() string {
	return "pkg/sentry/fs.DentAttr"
}

func (d *DentAttr) StateFields() []string {
	return []string{
		"Type",
		"InodeID",
	}
}

func (d *DentAttr) beforeSave() {}

func (d *DentAttr) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.Type)
	stateSinkObject.Save(1, &d.InodeID)
}

func (d *DentAttr) afterLoad() {}

func (d *DentAttr) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.Type)
	stateSourceObject.Load(1, &d.InodeID)
}

func (s *SortedDentryMap) StateTypeName() string {
	return "pkg/sentry/fs.SortedDentryMap"
}

func (s *SortedDentryMap) StateFields() []string {
	return []string{
		"names",
		"entries",
	}
}

func (s *SortedDentryMap) beforeSave() {}

func (s *SortedDentryMap) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.names)
	stateSinkObject.Save(1, &s.entries)
}

func (s *SortedDentryMap) afterLoad() {}

func (s *SortedDentryMap) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.names)
	stateSourceObject.Load(1, &s.entries)
}

func (d *Dirent) StateTypeName() string {
	return "pkg/sentry/fs.Dirent"
}

func (d *Dirent) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"userVisible",
		"Inode",
		"name",
		"parent",
		"deleted",
		"mounted",
		"children",
	}
}

func (d *Dirent) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	var childrenValue map[string]*Dirent = d.saveChildren()
	stateSinkObject.SaveValue(7, childrenValue)
	stateSinkObject.Save(0, &d.AtomicRefCount)
	stateSinkObject.Save(1, &d.userVisible)
	stateSinkObject.Save(2, &d.Inode)
	stateSinkObject.Save(3, &d.name)
	stateSinkObject.Save(4, &d.parent)
	stateSinkObject.Save(5, &d.deleted)
	stateSinkObject.Save(6, &d.mounted)
}

func (d *Dirent) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.AtomicRefCount)
	stateSourceObject.Load(1, &d.userVisible)
	stateSourceObject.Load(2, &d.Inode)
	stateSourceObject.Load(3, &d.name)
	stateSourceObject.Load(4, &d.parent)
	stateSourceObject.Load(5, &d.deleted)
	stateSourceObject.Load(6, &d.mounted)
	stateSourceObject.LoadValue(7, new(map[string]*Dirent), func(y interface{}) { d.loadChildren(y.(map[string]*Dirent)) })
	stateSourceObject.AfterLoad(d.afterLoad)
}

func (d *DirentCache) StateTypeName() string {
	return "pkg/sentry/fs.DirentCache"
}

func (d *DirentCache) StateFields() []string {
	return []string{
		"maxSize",
		"limit",
	}
}

func (d *DirentCache) beforeSave() {}

func (d *DirentCache) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	if !state.IsZeroValue(&d.currentSize) {
		state.Failf("currentSize is %#v, expected zero", &d.currentSize)
	}
	if !state.IsZeroValue(&d.list) {
		state.Failf("list is %#v, expected zero", &d.list)
	}
	stateSinkObject.Save(0, &d.maxSize)
	stateSinkObject.Save(1, &d.limit)
}

func (d *DirentCache) afterLoad() {}

func (d *DirentCache) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.maxSize)
	stateSourceObject.Load(1, &d.limit)
}

func (d *DirentCacheLimiter) StateTypeName() string {
	return "pkg/sentry/fs.DirentCacheLimiter"
}

func (d *DirentCacheLimiter) StateFields() []string {
	return []string{
		"max",
	}
}

func (d *DirentCacheLimiter) beforeSave() {}

func (d *DirentCacheLimiter) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	if !state.IsZeroValue(&d.count) {
		state.Failf("count is %#v, expected zero", &d.count)
	}
	stateSinkObject.Save(0, &d.max)
}

func (d *DirentCacheLimiter) afterLoad() {}

func (d *DirentCacheLimiter) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.max)
}

func (d *direntList) StateTypeName() string {
	return "pkg/sentry/fs.direntList"
}

func (d *direntList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (d *direntList) beforeSave() {}

func (d *direntList) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.head)
	stateSinkObject.Save(1, &d.tail)
}

func (d *direntList) afterLoad() {}

func (d *direntList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.head)
	stateSourceObject.Load(1, &d.tail)
}

func (d *direntEntry) StateTypeName() string {
	return "pkg/sentry/fs.direntEntry"
}

func (d *direntEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (d *direntEntry) beforeSave() {}

func (d *direntEntry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.next)
	stateSinkObject.Save(1, &d.prev)
}

func (d *direntEntry) afterLoad() {}

func (d *direntEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.next)
	stateSourceObject.Load(1, &d.prev)
}

func (e *eventList) StateTypeName() string {
	return "pkg/sentry/fs.eventList"
}

func (e *eventList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (e *eventList) beforeSave() {}

func (e *eventList) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.head)
	stateSinkObject.Save(1, &e.tail)
}

func (e *eventList) afterLoad() {}

func (e *eventList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.head)
	stateSourceObject.Load(1, &e.tail)
}

func (e *eventEntry) StateTypeName() string {
	return "pkg/sentry/fs.eventEntry"
}

func (e *eventEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *eventEntry) beforeSave() {}

func (e *eventEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *eventEntry) afterLoad() {}

func (e *eventEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (f *File) StateTypeName() string {
	return "pkg/sentry/fs.File"
}

func (f *File) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"UniqueID",
		"Dirent",
		"flags",
		"async",
		"FileOperations",
		"offset",
	}
}

func (f *File) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.AtomicRefCount)
	stateSinkObject.Save(1, &f.UniqueID)
	stateSinkObject.Save(2, &f.Dirent)
	stateSinkObject.Save(3, &f.flags)
	stateSinkObject.Save(4, &f.async)
	stateSinkObject.Save(5, &f.FileOperations)
	stateSinkObject.Save(6, &f.offset)
}

func (f *File) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.AtomicRefCount)
	stateSourceObject.Load(1, &f.UniqueID)
	stateSourceObject.Load(2, &f.Dirent)
	stateSourceObject.Load(3, &f.flags)
	stateSourceObject.Load(4, &f.async)
	stateSourceObject.LoadWait(5, &f.FileOperations)
	stateSourceObject.Load(6, &f.offset)
	stateSourceObject.AfterLoad(f.afterLoad)
}

func (o *overlayFileOperations) StateTypeName() string {
	return "pkg/sentry/fs.overlayFileOperations"
}

func (o *overlayFileOperations) StateFields() []string {
	return []string{
		"upper",
		"lower",
		"dirCursor",
	}
}

func (o *overlayFileOperations) beforeSave() {}

func (o *overlayFileOperations) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.upper)
	stateSinkObject.Save(1, &o.lower)
	stateSinkObject.Save(2, &o.dirCursor)
}

func (o *overlayFileOperations) afterLoad() {}

func (o *overlayFileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.upper)
	stateSourceObject.Load(1, &o.lower)
	stateSourceObject.Load(2, &o.dirCursor)
}

func (o *overlayMappingIdentity) StateTypeName() string {
	return "pkg/sentry/fs.overlayMappingIdentity"
}

func (o *overlayMappingIdentity) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"id",
		"overlayFile",
	}
}

func (o *overlayMappingIdentity) beforeSave() {}

func (o *overlayMappingIdentity) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.AtomicRefCount)
	stateSinkObject.Save(1, &o.id)
	stateSinkObject.Save(2, &o.overlayFile)
}

func (o *overlayMappingIdentity) afterLoad() {}

func (o *overlayMappingIdentity) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.AtomicRefCount)
	stateSourceObject.Load(1, &o.id)
	stateSourceObject.Load(2, &o.overlayFile)
}

func (m *MountSourceFlags) StateTypeName() string {
	return "pkg/sentry/fs.MountSourceFlags"
}

func (m *MountSourceFlags) StateFields() []string {
	return []string{
		"ReadOnly",
		"NoAtime",
		"ForcePageCache",
		"NoExec",
	}
}

func (m *MountSourceFlags) beforeSave() {}

func (m *MountSourceFlags) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.ReadOnly)
	stateSinkObject.Save(1, &m.NoAtime)
	stateSinkObject.Save(2, &m.ForcePageCache)
	stateSinkObject.Save(3, &m.NoExec)
}

func (m *MountSourceFlags) afterLoad() {}

func (m *MountSourceFlags) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.ReadOnly)
	stateSourceObject.Load(1, &m.NoAtime)
	stateSourceObject.Load(2, &m.ForcePageCache)
	stateSourceObject.Load(3, &m.NoExec)
}

func (f *FileFlags) StateTypeName() string {
	return "pkg/sentry/fs.FileFlags"
}

func (f *FileFlags) StateFields() []string {
	return []string{
		"Direct",
		"NonBlocking",
		"DSync",
		"Sync",
		"Append",
		"Read",
		"Write",
		"Pread",
		"Pwrite",
		"Directory",
		"Async",
		"LargeFile",
		"NonSeekable",
		"Truncate",
	}
}

func (f *FileFlags) beforeSave() {}

func (f *FileFlags) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Direct)
	stateSinkObject.Save(1, &f.NonBlocking)
	stateSinkObject.Save(2, &f.DSync)
	stateSinkObject.Save(3, &f.Sync)
	stateSinkObject.Save(4, &f.Append)
	stateSinkObject.Save(5, &f.Read)
	stateSinkObject.Save(6, &f.Write)
	stateSinkObject.Save(7, &f.Pread)
	stateSinkObject.Save(8, &f.Pwrite)
	stateSinkObject.Save(9, &f.Directory)
	stateSinkObject.Save(10, &f.Async)
	stateSinkObject.Save(11, &f.LargeFile)
	stateSinkObject.Save(12, &f.NonSeekable)
	stateSinkObject.Save(13, &f.Truncate)
}

func (f *FileFlags) afterLoad() {}

func (f *FileFlags) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Direct)
	stateSourceObject.Load(1, &f.NonBlocking)
	stateSourceObject.Load(2, &f.DSync)
	stateSourceObject.Load(3, &f.Sync)
	stateSourceObject.Load(4, &f.Append)
	stateSourceObject.Load(5, &f.Read)
	stateSourceObject.Load(6, &f.Write)
	stateSourceObject.Load(7, &f.Pread)
	stateSourceObject.Load(8, &f.Pwrite)
	stateSourceObject.Load(9, &f.Directory)
	stateSourceObject.Load(10, &f.Async)
	stateSourceObject.Load(11, &f.LargeFile)
	stateSourceObject.Load(12, &f.NonSeekable)
	stateSourceObject.Load(13, &f.Truncate)
}

func (i *Inode) StateTypeName() string {
	return "pkg/sentry/fs.Inode"
}

func (i *Inode) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"InodeOperations",
		"StableAttr",
		"LockCtx",
		"Watches",
		"MountSource",
		"overlay",
	}
}

func (i *Inode) beforeSave() {}

func (i *Inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.AtomicRefCount)
	stateSinkObject.Save(1, &i.InodeOperations)
	stateSinkObject.Save(2, &i.StableAttr)
	stateSinkObject.Save(3, &i.LockCtx)
	stateSinkObject.Save(4, &i.Watches)
	stateSinkObject.Save(5, &i.MountSource)
	stateSinkObject.Save(6, &i.overlay)
}

func (i *Inode) afterLoad() {}

func (i *Inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.AtomicRefCount)
	stateSourceObject.Load(1, &i.InodeOperations)
	stateSourceObject.Load(2, &i.StableAttr)
	stateSourceObject.Load(3, &i.LockCtx)
	stateSourceObject.Load(4, &i.Watches)
	stateSourceObject.Load(5, &i.MountSource)
	stateSourceObject.Load(6, &i.overlay)
}

func (l *LockCtx) StateTypeName() string {
	return "pkg/sentry/fs.LockCtx"
}

func (l *LockCtx) StateFields() []string {
	return []string{
		"Posix",
		"BSD",
	}
}

func (l *LockCtx) beforeSave() {}

func (l *LockCtx) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.Posix)
	stateSinkObject.Save(1, &l.BSD)
}

func (l *LockCtx) afterLoad() {}

func (l *LockCtx) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.Posix)
	stateSourceObject.Load(1, &l.BSD)
}

func (w *Watches) StateTypeName() string {
	return "pkg/sentry/fs.Watches"
}

func (w *Watches) StateFields() []string {
	return []string{
		"ws",
		"unlinked",
	}
}

func (w *Watches) beforeSave() {}

func (w *Watches) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.ws)
	stateSinkObject.Save(1, &w.unlinked)
}

func (w *Watches) afterLoad() {}

func (w *Watches) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.ws)
	stateSourceObject.Load(1, &w.unlinked)
}

func (i *Inotify) StateTypeName() string {
	return "pkg/sentry/fs.Inotify"
}

func (i *Inotify) StateFields() []string {
	return []string{
		"id",
		"events",
		"scratch",
		"nextWatch",
		"watches",
	}
}

func (i *Inotify) beforeSave() {}

func (i *Inotify) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.id)
	stateSinkObject.Save(1, &i.events)
	stateSinkObject.Save(2, &i.scratch)
	stateSinkObject.Save(3, &i.nextWatch)
	stateSinkObject.Save(4, &i.watches)
}

func (i *Inotify) afterLoad() {}

func (i *Inotify) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.id)
	stateSourceObject.Load(1, &i.events)
	stateSourceObject.Load(2, &i.scratch)
	stateSourceObject.Load(3, &i.nextWatch)
	stateSourceObject.Load(4, &i.watches)
}

func (e *Event) StateTypeName() string {
	return "pkg/sentry/fs.Event"
}

func (e *Event) StateFields() []string {
	return []string{
		"eventEntry",
		"wd",
		"mask",
		"cookie",
		"len",
		"name",
	}
}

func (e *Event) beforeSave() {}

func (e *Event) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.eventEntry)
	stateSinkObject.Save(1, &e.wd)
	stateSinkObject.Save(2, &e.mask)
	stateSinkObject.Save(3, &e.cookie)
	stateSinkObject.Save(4, &e.len)
	stateSinkObject.Save(5, &e.name)
}

func (e *Event) afterLoad() {}

func (e *Event) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.eventEntry)
	stateSourceObject.Load(1, &e.wd)
	stateSourceObject.Load(2, &e.mask)
	stateSourceObject.Load(3, &e.cookie)
	stateSourceObject.Load(4, &e.len)
	stateSourceObject.Load(5, &e.name)
}

func (w *Watch) StateTypeName() string {
	return "pkg/sentry/fs.Watch"
}

func (w *Watch) StateFields() []string {
	return []string{
		"owner",
		"wd",
		"target",
		"unpinned",
		"mask",
		"pins",
	}
}

func (w *Watch) beforeSave() {}

func (w *Watch) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.owner)
	stateSinkObject.Save(1, &w.wd)
	stateSinkObject.Save(2, &w.target)
	stateSinkObject.Save(3, &w.unpinned)
	stateSinkObject.Save(4, &w.mask)
	stateSinkObject.Save(5, &w.pins)
}

func (w *Watch) afterLoad() {}

func (w *Watch) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.owner)
	stateSourceObject.Load(1, &w.wd)
	stateSourceObject.Load(2, &w.target)
	stateSourceObject.Load(3, &w.unpinned)
	stateSourceObject.Load(4, &w.mask)
	stateSourceObject.Load(5, &w.pins)
}

func (m *MountSource) StateTypeName() string {
	return "pkg/sentry/fs.MountSource"
}

func (m *MountSource) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"MountSourceOperations",
		"FilesystemType",
		"Flags",
		"fscache",
		"direntRefs",
	}
}

func (m *MountSource) beforeSave() {}

func (m *MountSource) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.AtomicRefCount)
	stateSinkObject.Save(1, &m.MountSourceOperations)
	stateSinkObject.Save(2, &m.FilesystemType)
	stateSinkObject.Save(3, &m.Flags)
	stateSinkObject.Save(4, &m.fscache)
	stateSinkObject.Save(5, &m.direntRefs)
}

func (m *MountSource) afterLoad() {}

func (m *MountSource) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.AtomicRefCount)
	stateSourceObject.Load(1, &m.MountSourceOperations)
	stateSourceObject.Load(2, &m.FilesystemType)
	stateSourceObject.Load(3, &m.Flags)
	stateSourceObject.Load(4, &m.fscache)
	stateSourceObject.Load(5, &m.direntRefs)
}

func (s *SimpleMountSourceOperations) StateTypeName() string {
	return "pkg/sentry/fs.SimpleMountSourceOperations"
}

func (s *SimpleMountSourceOperations) StateFields() []string {
	return []string{
		"keep",
		"revalidate",
		"cacheReaddir",
	}
}

func (s *SimpleMountSourceOperations) beforeSave() {}

func (s *SimpleMountSourceOperations) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.keep)
	stateSinkObject.Save(1, &s.revalidate)
	stateSinkObject.Save(2, &s.cacheReaddir)
}

func (s *SimpleMountSourceOperations) afterLoad() {}

func (s *SimpleMountSourceOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.keep)
	stateSourceObject.Load(1, &s.revalidate)
	stateSourceObject.Load(2, &s.cacheReaddir)
}

func (o *overlayMountSourceOperations) StateTypeName() string {
	return "pkg/sentry/fs.overlayMountSourceOperations"
}

func (o *overlayMountSourceOperations) StateFields() []string {
	return []string{
		"upper",
		"lower",
	}
}

func (o *overlayMountSourceOperations) beforeSave() {}

func (o *overlayMountSourceOperations) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.upper)
	stateSinkObject.Save(1, &o.lower)
}

func (o *overlayMountSourceOperations) afterLoad() {}

func (o *overlayMountSourceOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.upper)
	stateSourceObject.Load(1, &o.lower)
}

func (o *overlayFilesystem) StateTypeName() string {
	return "pkg/sentry/fs.overlayFilesystem"
}

func (o *overlayFilesystem) StateFields() []string {
	return []string{}
}

func (o *overlayFilesystem) beforeSave() {}

func (o *overlayFilesystem) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
}

func (o *overlayFilesystem) afterLoad() {}

func (o *overlayFilesystem) StateLoad(stateSourceObject state.Source) {
}

func (m *Mount) StateTypeName() string {
	return "pkg/sentry/fs.Mount"
}

func (m *Mount) StateFields() []string {
	return []string{
		"ID",
		"ParentID",
		"root",
		"previous",
	}
}

func (m *Mount) beforeSave() {}

func (m *Mount) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.ID)
	stateSinkObject.Save(1, &m.ParentID)
	stateSinkObject.Save(2, &m.root)
	stateSinkObject.Save(3, &m.previous)
}

func (m *Mount) afterLoad() {}

func (m *Mount) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.ID)
	stateSourceObject.Load(1, &m.ParentID)
	stateSourceObject.Load(2, &m.root)
	stateSourceObject.Load(3, &m.previous)
}

func (m *MountNamespace) StateTypeName() string {
	return "pkg/sentry/fs.MountNamespace"
}

func (m *MountNamespace) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"userns",
		"root",
		"mounts",
		"mountID",
	}
}

func (m *MountNamespace) beforeSave() {}

func (m *MountNamespace) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.AtomicRefCount)
	stateSinkObject.Save(1, &m.userns)
	stateSinkObject.Save(2, &m.root)
	stateSinkObject.Save(3, &m.mounts)
	stateSinkObject.Save(4, &m.mountID)
}

func (m *MountNamespace) afterLoad() {}

func (m *MountNamespace) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.AtomicRefCount)
	stateSourceObject.Load(1, &m.userns)
	stateSourceObject.Load(2, &m.root)
	stateSourceObject.Load(3, &m.mounts)
	stateSourceObject.Load(4, &m.mountID)
}

func (o *overlayEntry) StateTypeName() string {
	return "pkg/sentry/fs.overlayEntry"
}

func (o *overlayEntry) StateFields() []string {
	return []string{
		"lowerExists",
		"lower",
		"mappings",
		"upper",
		"dirCache",
	}
}

func (o *overlayEntry) beforeSave() {}

func (o *overlayEntry) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.lowerExists)
	stateSinkObject.Save(1, &o.lower)
	stateSinkObject.Save(2, &o.mappings)
	stateSinkObject.Save(3, &o.upper)
	stateSinkObject.Save(4, &o.dirCache)
}

func (o *overlayEntry) afterLoad() {}

func (o *overlayEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.lowerExists)
	stateSourceObject.Load(1, &o.lower)
	stateSourceObject.Load(2, &o.mappings)
	stateSourceObject.Load(3, &o.upper)
	stateSourceObject.Load(4, &o.dirCache)
}

func init() {
	state.Register((*StableAttr)(nil))
	state.Register((*UnstableAttr)(nil))
	state.Register((*AttrMask)(nil))
	state.Register((*PermMask)(nil))
	state.Register((*FilePermissions)(nil))
	state.Register((*FileOwner)(nil))
	state.Register((*DentAttr)(nil))
	state.Register((*SortedDentryMap)(nil))
	state.Register((*Dirent)(nil))
	state.Register((*DirentCache)(nil))
	state.Register((*DirentCacheLimiter)(nil))
	state.Register((*direntList)(nil))
	state.Register((*direntEntry)(nil))
	state.Register((*eventList)(nil))
	state.Register((*eventEntry)(nil))
	state.Register((*File)(nil))
	state.Register((*overlayFileOperations)(nil))
	state.Register((*overlayMappingIdentity)(nil))
	state.Register((*MountSourceFlags)(nil))
	state.Register((*FileFlags)(nil))
	state.Register((*Inode)(nil))
	state.Register((*LockCtx)(nil))
	state.Register((*Watches)(nil))
	state.Register((*Inotify)(nil))
	state.Register((*Event)(nil))
	state.Register((*Watch)(nil))
	state.Register((*MountSource)(nil))
	state.Register((*SimpleMountSourceOperations)(nil))
	state.Register((*overlayMountSourceOperations)(nil))
	state.Register((*overlayFilesystem)(nil))
	state.Register((*Mount)(nil))
	state.Register((*MountNamespace)(nil))
	state.Register((*overlayEntry)(nil))
}
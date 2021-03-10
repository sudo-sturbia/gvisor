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

package icmp_bind_test

import (
	"context"
	"flag"
	"net"
	"testing"
	"time"

	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/test/packetimpact/testbench"
)

func init() {
	testbench.Initialize(flag.CommandLine)
	testbench.RPCTimeout = 500 * time.Millisecond
}

func TestICMPSocketBind(t *testing.T) {
	dut := testbench.NewDUT(t)

	tests := map[string]struct {
		bindTo          net.IP
		wantErr         unix.Errno
		wantErrNetstack *unix.Errno
	}{
		"IPv4Zero": {
			bindTo:  net.IPv4zero,
			wantErr: 0,
		},
		"IPv4Loopback": {
			bindTo:  net.IPv4(127, 0, 0, 1),
			wantErr: 0,
		},
		"IPv4Unicast": {
			bindTo:  dut.Net.RemoteIPv4,
			wantErr: 0,
		},
		"IPv4UnknownUnicast": {
			bindTo:  dut.Net.LocalIPv4,
			wantErr: unix.EADDRNOTAVAIL,
		},
		"IPv4MulticastAllSys": {
			bindTo:  net.IPv4allsys,
			wantErr: unix.EADDRNOTAVAIL,
		},
		"IPv4Broadcast": {
			bindTo:  net.IPv4bcast,
			wantErr: unix.EADDRNOTAVAIL,
			// TODO(gvisor.dev/issue/5711): Remove this field once ICMP sockets
			// are no longer allowed to bind to broadcast addresses.
			wantErrNetstack: newErrno(0),
		},
		"IPv4SubnetBroadcast": {
			bindTo:  dut.Net.SubnetBroadcast(),
			wantErr: unix.EADDRNOTAVAIL,
			// TODO(gvisor.dev/issue/5711): Remove this field once ICMP sockets
			// are no longer allowed to bind to broadcast addresses.
			wantErrNetstack: newErrno(0),
		},
		"IPv6Zero": {
			bindTo:  net.IPv6zero,
			wantErr: 0,
		},
		"IPv6Unicast": {
			bindTo:  dut.Net.RemoteIPv6,
			wantErr: 0,
		},
		"IPv6UnknownUnicast": {
			bindTo:  dut.Net.LocalIPv6,
			wantErr: unix.EADDRNOTAVAIL,
		},
		"IPv6MulticastInterfaceLocalAllNodes": {
			bindTo:  net.IPv6interfacelocalallnodes,
			wantErr: unix.EINVAL,
			// TODO(gvisor.dev/issue/5711): Remove this field once ICMPv6
			// sockets return EINVAL when binding to IPv6 multicast addresses.
			wantErrNetstack: newErrno(unix.EADDRNOTAVAIL),
		},
		"IPv6MulticastLinkLocalAllNodes": {
			bindTo:  net.IPv6linklocalallnodes,
			wantErr: unix.EINVAL,
			// TODO(gvisor.dev/issue/5711): Remove this field once ICMPv6
			// sockets return EINVAL when binding to IPv6 multicast addresses.
			wantErrNetstack: newErrno(unix.EADDRNOTAVAIL),
		},
		"IPv6MulticastLinkLocalAllRouters": {
			bindTo:  net.IPv6linklocalallrouters,
			wantErr: unix.EINVAL,
			// TODO(gvisor.dev/issue/5711): Remove this field once ICMPv6
			// sockets return EINVAL when binding to IPv6 multicast addresses.
			wantErrNetstack: newErrno(unix.EADDRNOTAVAIL),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var socketFD int32
			var sockaddr unix.Sockaddr

			if test.bindTo.To4() != nil {
				socketFD = dut.Socket(t, unix.AF_INET, unix.SOCK_DGRAM, unix.IPPROTO_ICMP)
				bindTo := unix.SockaddrInet4{}
				copy(bindTo.Addr[:], test.bindTo.To4())
				sockaddr = &bindTo
			} else {
				socketFD = dut.Socket(t, unix.AF_INET6, unix.SOCK_DGRAM, unix.IPPROTO_ICMPV6)
				bindTo := unix.SockaddrInet6{
					ZoneId: dut.Net.RemoteDevID,
				}
				copy(bindTo.Addr[:], test.bindTo.To16())
				sockaddr = &bindTo
			}

			ctx := context.Background()
			ret, err := dut.BindWithErrno(ctx, t, socketFD, sockaddr)

			wantRet, wantErr := int32(0), test.wantErr

			testingNetstack := !testbench.Native || dut.Uname.OperatingSystem == "Fuchsia"
			if test.wantErrNetstack != nil && testingNetstack {
				wantErr = *test.wantErrNetstack
			}

			if wantErr != 0 {
				wantRet = -1
			}

			if ret != wantRet {
				t.Errorf("got dut.BindWithErrno(_, _, %d, %s) = (%d, _), want (%d, _)", socketFD, test.bindTo, ret, wantRet)
			}
			if err != wantErr {
				t.Errorf("got dut.BindWithErrno(_, _, %d, %s) = (_, %s), want (_, %s)", socketFD, test.bindTo, err, wantErr)
			}
		})
	}
}

func newErrno(err unix.Errno) *unix.Errno {
	return &err
}

// automatically generated by stateify.

package signalfd

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *SignalFileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/signalfd.SignalFileDescription"
}

func (s *SignalFileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"target",
		"mask",
	}
}

func (s *SignalFileDescription) beforeSave() {}

func (s *SignalFileDescription) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.vfsfd)
	stateSinkObject.Save(1, &s.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &s.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &s.NoLockFD)
	stateSinkObject.Save(4, &s.target)
	stateSinkObject.Save(5, &s.mask)
}

func (s *SignalFileDescription) afterLoad() {}

func (s *SignalFileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.vfsfd)
	stateSourceObject.Load(1, &s.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &s.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &s.NoLockFD)
	stateSourceObject.Load(4, &s.target)
	stateSourceObject.Load(5, &s.mask)
}

func init() {
	state.Register((*SignalFileDescription)(nil))
}
package lscolors

import (
	"os"
)

type feature int

const (
	featureInvalid feature = iota

	featureOrphanedSymlink
	featureSymlink

	featureMultiHardLink

	featureNamedPipe
	featureSocket
	featureDoor
	featureBlockDevice
	featureCharDevice

	featureWorldWritableStickyDirectory
	featureWorldWritableDirectory
	featureStickyDirectory
	featureDirectory

	featureCapability

	featureSetuid
	featureSetgid
	featureExecutable

	featureRegular
)

// Some platforms, such as Windows, have simulated UNIX style permission masks.
// On Windows the only two permission masks are 0o666 (RW) and 0o444 (RO).
const worldWritable = 0o002

func determineFeature(fname string, mh bool) (feature, error) {
	_ = "STUB: not implemented"
	return *new(feature), nil
}

// Symlink and OrphanedSymlink has highest precedence

// featureMultiHardLink

// type bits features

// Never on Windows

// There is no dedicated os.Mode* flag for block device. On all
// supported Unix platforms, when os.ModeDevice is set but
// os.ModeCharDevice is not, the file is a block device (i.e.
// syscall.S_IFBLK is set). On Windows, this branch is unreachable.
//
// On Plan9, this in inaccurate.

// Perm bits features for directory

// TODO(xiaq): Support featureCapacity

// Perm bits features for regular files

// Check extension

func is(m, p os.FileMode) bool { _ = "STUB: not implemented"; return false }

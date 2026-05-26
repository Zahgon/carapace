package util

// TODO rename package update/optimize functions

// FindReverse traverses the filetree upwards to find given file/directory.
func FindReverse(path string, name string) (target string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HasPathPrefix checks if given string has a path prefix.
func HasPathPrefix(s string) bool { _ = "STUB: not implemented"; return false }

// HasVolumePrefix checks if given path has a volume prefix (only for GOOS=windows).
func HasVolumePrefix(s string) bool { _ = "STUB: not implemented"; return false }

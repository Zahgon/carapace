// Package cache provides cache keys
package key

// Key provides a cache key.
type Key func() (string, error)

// String creates a CacheKey for given strings.
func String(s ...string) Key { _ = "STUB: not implemented"; return *new(Key) }

// FileChecksum creates a CacheKey for given file.
func FileChecksum(file string) Key { _ = "STUB: not implemented"; return *new(Key) }

// FileStats creates a CacheKey for given file.
func FileStats(file string) Key { _ = "STUB: not implemented"; return *new(Key) }

func FolderStats(folder string) Key { _ = "STUB: not implemented"; return *new(Key) }

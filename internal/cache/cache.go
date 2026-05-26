// Package cache provides disk cache for Actions
package cache

import (
	"time"

	"github.com/carapace-sh/carapace/internal/export"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

func WriteE(file string, e export.Export) (err error) { _ = "STUB: not implemented"; return nil }

func Write(file string, content []byte) (err error) { _ = "STUB: not implemented"; return nil }

func LoadE(file string, timeout time.Duration) (*export.Export, error) {
	_ = "STUB: not implemented" // TODO reference
	return nil, nil
}

func Load(file string, timeout time.Duration) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CacheDir creates a cache folder for current user and returns the path.
func CacheDir(name string) (dir string, err error) { _ = "STUB: not implemented"; return "", nil }

// File returns the cache filename for given values
// TODO cleanup
func File(callerFile string, callerLine int, keys ...key.Key) (file string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func uidKeys(keys ...string) string { _ = "STUB: not implemented"; return "" }

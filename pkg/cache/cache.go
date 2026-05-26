package cache

import (
	"time"

	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// Cache caches a function for given duration and keys.
func Cache(timeout time.Duration, keys ...key.Key) func(f func() ([]byte, error)) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

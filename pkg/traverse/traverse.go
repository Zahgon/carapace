package traverse

import (
	"github.com/spf13/pflag"
)

type Context interface {
	Abs(s string) (string, error)
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
}

// Parent returns the first parent directory containing any of the given names/directories.
func Parent(names ...string) func(tc Context) (string, error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO also stop at `~`
func traverse(path string, name string) (target string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Flag returns the value of given flag.
func Flag(f *pflag.Flag) func(tc Context) (string, error) { _ = "STUB: not implemented"; return nil }

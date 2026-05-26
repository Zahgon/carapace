// Package uid provides unique identifiers
package uid

import (
	"net/url"
	"sync"

	"github.com/carapace-sh/carapace/internal/pflagfork"
	"github.com/spf13/cobra"
)

type Context interface {
	Abs(s string) (string, error)
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
}

// UidF TODO experimental
func UidF(scheme, host string, opts ...string) func(v string, uc Context) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil
}

// implicitly skip empty values

// Command creates a uid for given command.
func Command(cmd *cobra.Command) *url.URL { _ = "STUB: not implemented"; return nil }

// TODO slices.Reverse

// reverse reverses the elements of the slice in place.
func reverse(s []string) { _ = "STUB: not implemented"; return }

var mLocalFlags sync.Mutex

// Flag creates a uid for given flag.
func Flag(cmd *cobra.Command, flag *pflagfork.Flag) *url.URL { _ = "STUB: not implemented"; return nil }

func flagRecursive(cmd *cobra.Command, flag *pflagfork.Flag) *url.URL {
	_ = "STUB: not implemented"
	// Force flag merge; not thread-safe internally
	return nil
}

// Executable returns the name of the executable.
func Executable() string { _ = "STUB: not implemented"; return "" }

// safe fallback that should never happen

// for `go test -v ./...`

// alpine container workaround (gcompat)

// Map maps values to uids to simplify testing.
//
//	Map(
//	    "go.mod", "file://path/to/go.mod",
//	    "go.sum", "file://path/to/go.sum",
//	)
func Map(uids ...string) func(s string) (*url.URL, error) { _ = "STUB: not implemented"; return nil }

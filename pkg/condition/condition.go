package condition

import (
	"github.com/carapace-sh/carapace"
)

// Arch returns true if any of the given strings matches `runtime.GOARCH`.
func Arch(s ...string) func(c carapace.Context) bool { _ = "STUB: not implemented"; return nil }

// Arch returns true if any of the given strings matches `runtime.GOOS`.
func Os(s ...string) func(c carapace.Context) bool { _ = "STUB: not implemented"; return nil }

// Excutable returns true if any of the given strings matches an executable in PATH.
func Executable(s ...string) func(c carapace.Context) bool { _ = "STUB: not implemented"; return nil }

// TODO needs to use Context.Env

// Retuns true if given string is a valid file or directory.
func File(s string) func(c carapace.Context) bool { _ = "STUB: not implemented"; return nil }

// CompletingPath returns true when `Context.Value` has a path prefix.
func CompletingPath(c carapace.Context) bool { _ = "STUB: not implemented"; return false }

// CompletingPathS is like CompletingPathS but also checks for path separator `/`
func CompletingPathS(c carapace.Context) bool { _ = "STUB: not implemented"; return false }

// TODO support windows backslash at some point?

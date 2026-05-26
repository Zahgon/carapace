// Package carapace is a command argument completion generator for spf13/cobra
package carapace

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Carapace wraps cobra.Command to define completions.
type Carapace struct {
	cmd *cobra.Command
}

// Gen initialized Carapace for given command.
func Gen(cmd *cobra.Command) *Carapace { _ = "STUB: not implemented"; return nil }

// PreRun sets a function to be run before completion.
func (c Carapace) PreRun(f func(cmd *cobra.Command, args []string)) {
	_ = "STUB: not implemented"
	return
}

// TODO yuck - probably best to append to a slice in storage

// PreInvoke sets a function to alter actions before they are invoked.
func (c Carapace) PreInvoke(f func(cmd *cobra.Command, flag *pflag.Flag, action Action) Action) {
	_ = "STUB: not implemented"
	return
}

// PositionalCompletion defines completion for positional arguments using a list of Actions.
func (c Carapace) PositionalCompletion(action ...Action) { _ = "STUB: not implemented"; return }

// PositionalAnyCompletion defines completion for any positional arguments not already defined.
func (c Carapace) PositionalAnyCompletion(action Action) { _ = "STUB: not implemented"; return }

// DashCompletion defines completion for positional arguments after dash (`--`) using a list of Actions.
func (c Carapace) DashCompletion(action ...Action) { _ = "STUB: not implemented"; return }

// DashAnyCompletion defines completion for any positional arguments after dash (`--`) not already defined.
func (c Carapace) DashAnyCompletion(action Action) { _ = "STUB: not implemented"; return }

// FlagCompletion defines completion for flags using a map consisting of name and Action.
func (c Carapace) FlagCompletion(actions ActionMap) { _ = "STUB: not implemented"; return }

const annotation_standalone = "carapace_standalone"

// Standalone prevents cobra defaults interfering with standalone mode (e.g. implicit help command).
func (c Carapace) Standalone() { _ = "STUB: not implemented"; return }

// Snippet creates completion script for given shell.
func (c Carapace) Snippet(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// IsCallback returns true if current program invocation is a callback.
func IsCallback() bool { _ = "STUB: not implemented"; return false }

// Test verifies the configuration (e.g. flag name exists)
//
//	func TestCarapace(t *testing.T) {
//	    carapace.Test(t)
//	}
func Test(t interface{ Error(args ...any) }) { _ = "STUB: not implemented"; return }

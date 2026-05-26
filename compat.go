package carapace

import (
	"github.com/spf13/cobra"
)

func registerValidArgsFunction(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// TODO just IvokedAction{} ok?

func registerFlagCompletion(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// skip if not defined in carapace

// skip if already defined in cobra

// TODO cmd might differ for persistentflags and either way args or cmd will be wrong

func cobraValuesFor(action InvokedAction) []string { _ = "STUB: not implemented"; return nil }

func cobraDirectiveFor(action InvokedAction) cobra.ShellCompDirective {
	_ = "STUB: not implemented"
	return *new(cobra.ShellCompDirective)
}

type compDirective cobra.ShellCompDirective

func (d compDirective) matches(cobraDirective cobra.ShellCompDirective) bool {
	_ = "STUB: not implemented"
	return false
}

func (d compDirective) ToA(values ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

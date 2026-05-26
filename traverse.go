package carapace

import (
	"github.com/spf13/cobra"
)

func traverse(cmd *cobra.Command, args []string) (Action, Context) {
	_ = "STUB: not implemented"
	return *new(Action), *new(Context)
}

// args consumed by current command
// positionals consumed by current command
// last encountered flag that still expects arguments
// TODO force  c.mergePersistentFlags() which is missing from c.Flags()

// dash encountered

// flag argument

// no more args expected

// dash

// flag

// subcommand

// positional

// skip looking for flags in dash arguments

// TODO shorthand series isn't correct anymore (can have value attached)
// TODO not aways correct

// TODO && len(context.Value) > 2 {
// TODO check if empty prefix

// TODO duplicated code

// dash argument

// flag argument

// flag

//nolint:govet

// positional or subcommand

func subcommand(cmd *cobra.Command, arg string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

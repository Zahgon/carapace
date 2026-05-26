package carapace

import (
	"github.com/spf13/cobra"
)

func actionPath(fileSuffixes []string, dirOnly bool) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// TODO should be fixed in Abs or wherever this is happening
// prevent `C:` -> `C:.`

// TODO colorist not returning the symlink color

func actionFlags(cmd *cobra.Command) Action { _ = "STUB: not implemented"; return *new(Action) }

// skip hidden flags

// skip deprecated flags

// don't repeat flag

// skip flag of group already set

// abort shorthand flag series if a previous one is not bool or count and requires an argument (no default value)

// multiparts completion for flags grouped with `.`

func initHelpCompletion(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

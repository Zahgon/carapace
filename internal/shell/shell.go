package shell

import (
	"github.com/carapace-sh/carapace/internal/common"
	"github.com/spf13/cobra"
)

// Snippet creates completion script for given shell.
func Snippet(cmd *cobra.Command, shell string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func Value(shell string, value string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented" // TODO use context instead?
	return ""
}

// explicit
// implicit for classic zsh side-by-side view

// shells with support for showing messages

// re-filter after clearance

func mergeFlags(values common.RawValues) { _ = "STUB: not implemented"; return }

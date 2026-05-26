package carapace

import (
	"github.com/spf13/cobra"
)

func addCompletionCommand(targetCmd *cobra.Command) { _ = "STUB: not implemented"; return }

// this should never happen

// TODO how to handle an explicit `_carapace` command?
// don't complete local `_carapace` in standalone mode

// TODO does not work with sandbox tests for `example _carapace ...`

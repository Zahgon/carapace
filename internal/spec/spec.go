// Package spec provides spec file generation for use with carapace-bin
package spec

import (
	"github.com/spf13/cobra"
)

// Spec generates the spec file.
func Spec(cmd *cobra.Command) string { _ = "STUB: not implemented"; return "" }

func command(cmd *cobra.Command) Command { _ = "STUB: not implemented"; return *new(Command) }

// TODO mutually exclusive flags

package common

import (
	"github.com/spf13/cobra"
)

type Group struct {
	Cmd *cobra.Command
}

func (g Group) Tag() string { _ = "STUB: not implemented"; return "" }

func (g Group) Style() string { _ = "STUB: not implemented"; return "" }

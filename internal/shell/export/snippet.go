// Package export provides command structure export
package export

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type command struct {
	Name            string
	Short           string
	Long            string    `json:",omitempty"`
	Aliases         []string  `json:",omitempty"`
	Commands        []command `json:",omitempty"`
	LocalFlags      []flag    `json:",omitempty"`
	PersistentFlags []flag    `json:",omitempty"`
}

type flag struct {
	Longhand    string `json:",omitempty"`
	Shorthand   string `json:",omitempty"`
	Usage       string
	Type        string
	NoOptDefVal string `json:",omitempty"`
}

func convertFlag(f *pflag.Flag) flag { _ = "STUB: not implemented"; return *new(flag) }

func convert(cmd *cobra.Command) command { _ = "STUB: not implemented"; return *new(command) }

// Snippet exports the command structure as json.
func Snippet(cmd *cobra.Command) string { _ = "STUB: not implemented"; return "" }

package zsh

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/env"
)

type namedDirectories map[string]string

// NamedDirectories provides rudimentary named directory support as these aren't expanded by zsh in the `${words}` provided to the compdef function
var NamedDirectories = make(namedDirectories)

func (nd *namedDirectories) match(s string) string { _ = "STUB: not implemented"; return "" }

// Matches checks if given string has a known named directory prefix
func (nd *namedDirectories) Matches(s string) bool { _ = "STUB: not implemented"; return false }

// Replace replaces a known named directory prefix with the actual folder
func (nd *namedDirectories) Replace(s string) string { _ = "STUB: not implemented"; return "" }

func init() {
	if hashDirs := env.Hashdirs(); hashDirs != "" {
		for line := range strings.SplitSeq(hashDirs, "\n") {
			if splitted := strings.SplitN(line, "=", 2); len(splitted) == 2 {
				NamedDirectories[splitted[0]] = splitted[1]
			}
		}
	}
}

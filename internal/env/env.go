package env

import (
	"github.com/carapace-sh/carapace/internal/mock"
)

const (
	CARAPACE_COLOR              = "CARAPACE_COLOR"              // enable color
	CARAPACE_COMPLINE           = "CARAPACE_COMPLINE"           // TODO
	CARAPACE_COVERDIR           = "CARAPACE_COVERDIR"           // coverage directory for sandbox tests
	CARAPACE_DESCRIPTION_LENGTH = "CARAPACE_DESCRIPTION_LENGTH" // maximum description length
	CARAPACE_EXPERIMENTAL       = "CARAPACE_EXPERIMENTAL"       // enable experimental features
	CARAPACE_HIDDEN             = "CARAPACE_HIDDEN"             // show hidden commands/flags
	CARAPACE_LENIENT            = "CARAPACE_LENIENT"            // allow unknown flags
	CARAPACE_LOG                = "CARAPACE_LOG"                // enable logging
	CARAPACE_MATCH              = "CARAPACE_MATCH"              // match case insensitive
	CARAPACE_MERGEFLAGS         = "CARAPACE_MERGEFLAGS"         // merge flags to single tag group
	CARAPACE_NOSPACE            = "CARAPACE_NOSPACE"            // nospace suffixes
	CARAPACE_SANDBOX            = "CARAPACE_SANDBOX"            // mock context for sandbox tests
	CARAPACE_TOOLTIP            = "CARAPACE_TOOLTIP"            // enable tooltip style
	CARAPACE_UNFILTERED         = "CARAPACE_UNFILTERED"         // skip the final filtering step
	CARAPACE_ZSH_HASH_DIRS      = "CARAPACE_ZSH_HASH_DIRS"      // zsh hash directories
	CLICOLOR                    = "CLICOLOR"                    // disable color
	NO_COLOR                    = "NO_COLOR"                    // disable color
)

func ColorDisabled() bool { _ = "STUB: not implemented"; return false }

// TODO multiple modes

func Experimental() bool { _ = "STUB: not implemented"; return false }

func Lenient() bool { _ = "STUB: not implemented"; return false }

func Hashdirs() string { _ = "STUB: not implemented"; return "" }

func Sandbox() (m *mock.Mock, err error) { _ = "STUB: not implemented"; return nil, nil }

func Log() bool { _ = "STUB: not implemented"; return false }

type hidden int

const (
	HIDDEN_NONE hidden = iota
	HIDDEN_EXCLUDE_CARAPACE
	HIDDEN_INCLUDE_CARAPACE
)

func Hidden() hidden { _ = "STUB: not implemented"; return *new(hidden) }

// 0 or error

func CoverDir() string { _ = "STUB: not implemented"; return "" }

// custom env for GOCOVERDIR so that it works together with `-coverprofile`

func isGoRun() bool {
	_ = "STUB: not implemented"
	// go 1.24+: /home/rsteube/.cache/go-build/a7/a7883331b606dc3250005f1abdf4d17dea05c9eccf4ece0df5518311a118d211-d/example
	// go 1.23-: /tmp/go-build1035720725/b001/exe/example
	return false
}

func Match() string {
	_ = "STUB: not implemented" // see match.Match
	return ""
}

func MergeFlags() (bool, bool) { _ = "STUB: not implemented"; return false, false }

func Nospace() string { _ = "STUB: not implemented"; return "" }

func Tooltip() bool { _ = "STUB: not implemented"; return false }

func Compline() string { _ = "STUB: not implemented"; return "" }

func Unfiltered() bool { _ = "STUB: not implemented"; return false }

func getBool(s string) bool { _ = "STUB: not implemented"; return false }

func DescriptionLength() int { _ = "STUB: not implemented"; return 0 }

package xonsh

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

var sanitizer = strings.NewReplacer( // TODO
	"\n", ``,
	"\t", ``,
	`'`, `\'`,
)

type richCompletion struct {
	Value       string
	Display     string
	Description string
	Style       string
}

// ActionRawValues formats values for xonsh.
func ActionRawValues(currentWord string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented"
	return ""
}

// backslash needs raw string

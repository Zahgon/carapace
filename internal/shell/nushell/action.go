package nushell

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

type record struct {
	Value       string        `json:"value"`
	Display     string        `json:"display"`
	Description string        `json:"description,omitempty"`
	Style       *nushellStyle `json:"style,omitempty"`
}

var sanitizer = strings.NewReplacer(
	"\n", ``,
	"\r", ``,
)

var escaper = strings.NewReplacer(
	`\`, `\\`,
	`"`, `\"`,
)

func sanitize(values []common.RawValue) []common.RawValue { _ = "STUB: not implemented"; return nil }

// ActionRawValues formats values for nushell.
func ActionRawValues(currentWord string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented"
	return ""
}

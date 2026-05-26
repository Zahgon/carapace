package bash

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

var sanitizer = strings.NewReplacer(
	"\n", ``,
	"\r", ``,
	"\t", ``,
)

var escapingQuotedReplacer = strings.NewReplacer(
	`\`, `\\`,
	`"`, `\"`,
	`$`, `\$`,
	"`", "\\`",
)

var escapingReplacer = strings.NewReplacer(
	`\`, `\\`,
	`&`, `\&`,
	`<`, `\<`,
	`>`, `\>`,
	"`", "\\`",
	`'`, `\'`,
	`"`, `\"`,
	`{`, `\{`,
	`}`, `\}`,
	`$`, `\$`,
	`#`, `\#`,
	`|`, `\|`,
	`?`, `\?`,
	`(`, `\(`,
	`)`, `\)`,
	`;`, `\;`,
	` `, `\ `,
	`[`, `\[`,
	`]`, `\]`,
	`*`, `\*`,
)

var displayReplacer = strings.NewReplacer(
	`${`, `\\\${`,
)

func commonPrefix(a, b string) string { _ = "STUB: not implemented"; return "" }

func commonDisplayPrefix(values ...common.RawValue) (prefix string) {
	_ = "STUB: not implemented"
	return ""
}

func commonValuePrefix(values ...common.RawValue) (prefix string) {
	_ = "STUB: not implemented"
	return ""
}

// ActionRawValues formats values for bash.
func ActionRawValues(currentWord string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented"
	return ""
}

// last segment of currentWord split by COMP_WORDBREAKS

// When all display values have the same prefix bash will insert is as partial completion (which skips prefixes/formatting).

// replace values with common value prefix

// prevent insertion of partial display values by prefixing one with space

// assume homedir expansion

func requiresQuoting(s string) bool { _ = "STUB: not implemented"; return false }

package tcsh

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

var sanitizer = strings.NewReplacer(
	"\n", ``,
	"\r", ``,
	"\t", ``,
)

var quoter = strings.NewReplacer(
	`&`, `\&`,
	`<`, `\<`,
	`>`, `\>`,
	"`", "\\`",
	`'`, `\'`,
	`"`, `\"`,
	`{`, ``, // TODO seems escaping is not working
	`}`, ``, // TODO seems escaping is not working
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
	`\`, `\\`,
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
	// last segment of currentWord split by COMP_WORDBREAKS
}

// TODO optimize

// When all display values have the same prefix bash will insert is as partial completion (which skips prefixes/formatting).

// replace values with common value prefix (`\001` is removed in snippet and compopt nospace will be set)
// TODO nospaceIndicator
//values = common.RawValuesFrom(commonValuePrefix(values...) + nospaceIndicator)

// prevent insertion of partial display values by prefixing one with space

// TODO seems actual value needs to be used or it won't be shown if the prefix doesn't match

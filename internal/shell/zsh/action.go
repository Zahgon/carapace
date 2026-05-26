package zsh

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

var sanitizer = strings.NewReplacer(
	"\n", ``,
	"\r", ``,
	"\t", ``,
)

var quotingReplacer = strings.NewReplacer(
	`'`, `'\''`,
)

var quotingEscapingReplacer = strings.NewReplacer(
	`\`, `\\`,
	`"`, `\"`,
	`$`, `\$`,
	"`", "\\`",
)

var defaultReplacer = strings.NewReplacer(
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
	`~`, `\~`,
)

// additional replacement for use with `_describe` in shell script
var describeReplacer = strings.NewReplacer(
	`\`, `\\`,
	`:`, `\:`,
)

func quoteValue(s string) string { _ = "STUB: not implemented"; return "" }

// assume file path expansion

type state int

const (
	DEFAULT_STATE state = iota
	// Word starts with `"`.
	// Values need to end with `"` as well.
	// Weirdly regardless whether there are additional quotes within the word.
	QUOTING_ESCAPING_STATE
	// Word starts with `'`.
	// Values need to end with `'` as well.
	// Weirdly regardless whether there are additional quotes within the word.
	QUOTING_STATE
	// Word starts and ends with `"`.
	// Space suffix somehow ends up within the quotes.
	//    `"action"<TAB>`
	//    `"action "<CURSOR>`
	// Workaround for now is to force nospace.
	FULL_QUOTING_ESCAPING_STATE
	// Word starts and ends with `'`.
	// Space suffix somehow ends up within the quotes.
	//    `'action'<TAB>`
	//    `'action '<CURSOR>`
	// Workaround for now is to force nospace.
	FULL_QUOTING_STATE
)

// ActionRawValues formats values for zsh
func ActionRawValues(currentWord string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO use token state to determine actual state (might have mixture).

// nospace workaround

// TODO check if this needs to be applied to description as well

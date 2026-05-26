package powershell

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

var sanitizer = strings.NewReplacer( // TODO
	"\n", ``,
	"\t", ``,
)

type completionResult struct {
	CompletionText string
	ListItemText   string
	ToolTip        string
}

// CompletionResult doesn't like empty parameters, so just replace with space if needed.
func ensureNotEmpty(s string) string { _ = "STUB: not implemented"; return "" }

// ActionRawValues formats values for powershell.
func ActionRawValues(currentWord string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented"
	return ""
}

// must not be empty - any empty `''` parameter in CompletionResult causes an error

func sgr(s string) string { _ = "STUB: not implemented"; return "" }

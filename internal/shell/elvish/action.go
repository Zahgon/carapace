package elvish

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

var sanitizer = strings.NewReplacer(
	"\n", ``,
	"\r", ``,
	"\t", ``,
)

func sanitize(values []common.RawValue) []common.RawValue { _ = "STUB: not implemented"; return nil }

type completion struct {
	Usage            string
	Messages         common.Messages
	DescriptionStyle string
	Candidates       []complexCandidate
}

type complexCandidate struct {
	Value       string
	Display     string
	Description string
	CodeSuffix  string
	Style       string
}

// ActionRawValues formats values for elvish.
func ActionRawValues(currentWord string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO edit:notify is persistent, so avoid spamming the user for now

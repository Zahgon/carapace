package cmd_clink

import (
	"strings"

	"github.com/carapace-sh/carapace/internal/common"
)

var sanitizer = strings.NewReplacer(
	"\n", ``,
	"\r", ``,
	"\t", ``,
)

func ActionRawValues(currentWord string, meta common.Meta, values common.RawValues) string {
	_ = "STUB: not implemented"
	return ""
}

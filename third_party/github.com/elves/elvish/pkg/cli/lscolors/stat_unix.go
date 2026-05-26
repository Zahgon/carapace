//go:build !windows && !plan9

package lscolors

import (
	"os"
)

func isMultiHardlink(info os.FileInfo) bool { _ = "STUB: not implemented"; return false }

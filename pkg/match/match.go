package match

import (
	"os"
	"strconv"
)

type Match int

const (
	CASE_SENSITIVE Match = iota
	CASE_INSENSITIVE
)

func (m Match) Equal(s, t string) bool { _ = "STUB: not implemented"; return false }

func (m Match) HasPrefix(s, prefix string) bool { _ = "STUB: not implemented"; return false }

func (m Match) TrimPrefix(s, prefix string) string { _ = "STUB: not implemented"; return "" }

var match = CASE_SENSITIVE

func init() {
	switch os.Getenv("CARAPACE_MATCH") {
	case "CASE_INSENSITIVE", strconv.Itoa(int(CASE_INSENSITIVE)):
		match = CASE_INSENSITIVE
	}
}

func Equal(s, t string) bool { _ = "STUB: not implemented"; return false }

func HasPrefix(s, prefix string) bool { _ = "STUB: not implemented"; return false }

func TrimPrefix(s, prefix string) string { _ = "STUB: not implemented"; return "" }

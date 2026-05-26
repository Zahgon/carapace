package pflagfork

import (
	"github.com/spf13/pflag"
)

// mode defines how flags are represented.
type mode int

const (
	Default         mode = iota // default behaviour
	ShorthandOnly               // only the shorthand should be used
	NameAsShorthand             // non-posix mode where the name is also added as shorthand (single `-` prefix)
)

type Flag struct {
	*pflag.Flag
	Prefix string
	Args   []string
}

func (f Flag) Nargs() int { _ = "STUB: not implemented"; return 0 }

func (f Flag) Mode() mode { _ = "STUB: not implemented"; return *new(mode) }

func (f Flag) OptargDelimiter() rune { _ = "STUB: not implemented"; return 0 }

func (f Flag) IsRepeatable() bool { _ = "STUB: not implemented"; return false }

func (f Flag) TakesValue() bool { _ = "STUB: not implemented"; return false }

func (f Flag) IsOptarg() bool { _ = "STUB: not implemented"; return false }

func (f Flag) Style() string { _ = "STUB: not implemented"; return "" }

func (f Flag) Required() bool { _ = "STUB: not implemented"; return false }

func (f Flag) Definition() string { _ = "STUB: not implemented"; return "" }

func (f Flag) Consumes(arg string) bool { _ = "STUB: not implemented"; return false }

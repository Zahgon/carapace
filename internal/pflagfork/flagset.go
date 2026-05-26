package pflagfork

import (
	"github.com/spf13/pflag"
)

type FlagSet struct {
	*pflag.FlagSet
}

func (f FlagSet) IsInterspersed() bool { _ = "STUB: not implemented"; return false }

func (f FlagSet) IsPosix() bool { _ = "STUB: not implemented"; return false }

func (f FlagSet) IsShorthandSeries(arg string) bool { _ = "STUB: not implemented"; return false }

func (f FlagSet) IsMutuallyExclusive(flag *pflag.Flag) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *FlagSet) VisitAll(fn func(*Flag)) { _ = "STUB: not implemented"; return }

func (fs FlagSet) LookupArg(arg string) (result *Flag) { _ = "STUB: not implemented"; return nil }

// In non-posix mode, try longhand first (single dash with name)
// to handle cases where name overlaps with shorthand

func (fs FlagSet) ShorthandLookup(name string) *Flag { _ = "STUB: not implemented"; return nil }

func (fs FlagSet) lookupPosixLonghandArg(arg string) (flag *Flag) {
	_ = "STUB: not implemented"
	return nil
}

// TODO needs to be sorted to try longest matching first

func (fs FlagSet) lookupPosixShorthandArg(arg string) *Flag { _ = "STUB: not implemented"; return nil }

func (fs FlagSet) lookupNonPosixShorthandArg(arg string) (result *Flag) {
	_ = "STUB: not implemented" // TODO pretty much duplicates longhand lookup
	return nil
}

// TODO needs to be sorted to try longest matching first

// LookupNonPosixLonghandArg looks up a non-posix longhand argument (single dash with name)
func (fs FlagSet) LookupNonPosixLonghandArg(arg string) *Flag {
	_ = "STUB: not implemented"
	return nil
}

// remove single dash

package carapace

import (
	"net/url"
	"time"

	"github.com/carapace-sh/carapace/internal/common"
	"github.com/carapace-sh/carapace/pkg/cache/key"
	"github.com/carapace-sh/carapace/pkg/style"
	pkgtraverse "github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/pkg/uid"
)

// Action indicates how to complete a flag or positional argument.
type Action struct {
	meta      common.Meta
	rawValues common.RawValues
	callback  CompletionCallback
}

// ActionMap maps Actions to an identifier.
type ActionMap map[string]Action

// CompletionCallback is executed during completion of associated flag or positional argument.
type CompletionCallback func(c Context) Action

// Cache cashes values of a CompletionCallback for given duration and keys.
func (a Action) Cache(timeout time.Duration, keys ...key.Key) Action {
	_ = "STUB: not implemented"
	return *
	// only relevant for callback actions
	new(Action)
}

// generate uid from wherever Cache() was called

// regenerate as cache keys might have changed due to invocation

// Chdir changes the current working directory to the named directory for the duration of invocation.
func (a Action) Chdir(dir string) Action { _ = "STUB: not implemented"; return *new(Action) }

// ChdirF is like Chdir but uses a function.
func (a Action) ChdirF(f func(tc pkgtraverse.Context) (string, error)) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// Filter filters given values.
//
//	carapace.ActionValues("A", "B", "C").Filter("B") // ["A", "C"]
func (a Action) Filter(values ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// FilterArgs filters Context.Args.
func (a Action) FilterArgs() Action { _ = "STUB: not implemented"; return *new(Action) }

// FilterArgs filters Context.Parts.
func (a Action) FilterParts() Action { _ = "STUB: not implemented"; return *new(Action) }

// Invoke executes the callback of an action if it exists (supports nesting).
func (a Action) Invoke(c Context) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

// List wraps the Action in an ActionMultiParts with given divider.
func (a Action) List(divider string) Action { _ = "STUB: not implemented"; return *new(Action) }

// MultiParts splits values of an Action by given dividers and completes each segment separately.
func (a Action) MultiParts(dividers ...string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// MultiPartsP is like MultiParts but with placeholders.
func (a Action) MultiPartsP(delimiter string, pattern string, f func(placeholder string, matches map[string]string) Action) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// skip this path as it doesn't match and is not a placeholder

// store entered data for placeholder (overwrite if duplicate)

// static segment matches so placeholders should be ignored for this index

// skip path as it is shorter than what was entered (must be after staticMatches being set)

// skip this path as it has a placeholder where a static segment was matched

// store segment as path matched so far and this is currently being completed

// TODO tag,..

// NoSpace disables space suffix for given characters (or all if none are given).
func (a Action) NoSpace(suffixes ...rune) Action { _ = "STUB: not implemented"; return *new(Action) }

// Prefix adds a prefix to values (only the ones inserted, not the display values).
//
//	carapace.ActionValues("melon", "drop", "fall").Prefix("water")
func (a Action) Prefix(prefix string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Retain retains given values.
//
//	carapace.ActionValues("A", "B", "C").Retain("A", "C") // ["A", "C"]
func (a Action) Retain(values ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Shift shifts positional arguments left `n` times.
func (a Action) Shift(n int) Action { _ = "STUB: not implemented"; return *new(Action) }

// Split splits `Context.Value` lexicographically and replaces `Context.Args` with the tokens.
func (a Action) Split() Action {
	_ = "STUB: not implemented"
	return *

	// SplitP is like Split but supports pipelines.
	new(Action)
}

func (a Action) SplitP() Action { _ = "STUB: not implemented"; return *new(Action) }

func (a Action) split(pipelines bool) Action { _ = "STUB: not implemented"; return *new(Action) }

// support redirects

// TODO special characters

// Style sets the style.
//
//	ActionValues("yes").Style(style.Green)
//	ActionValues("no").Style(style.Red)
func (a Action) Style(s string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Style sets the style using a function.
//
//	ActionValues("dir/", "test.txt").StyleF(style.ForPathExt)
//	ActionValues("true", "false").StyleF(style.ForKeyword)
func (a Action) StyleF(f func(s string, sc style.Context) string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// Style sets the style using a reference.
//
//	ActionValues("value").StyleR(&style.Carapace.Value)
//	ActionValues("description").StyleR(&style.Carapace.Value)
func (a Action) StyleR(s *string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Suffix adds a suffx to values (only the ones inserted, not the display values).
//
//	carapace.ActionValues("apple", "melon", "orange").Suffix("juice")
func (a Action) Suffix(suffix string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Suppress suppresses specific error messages using regular expressions.
func (a Action) Suppress(expr ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Tag sets the tag.
//
//	ActionValues("192.168.1.1", "127.0.0.1").Tag("interfaces").
func (a Action) Tag(tag string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Tag sets the tag using a function.
//
//	ActionValues("192.168.1.1", "127.0.0.1").TagF(func(value string) string {
//		return "interfaces"
//	})
func (a Action) TagF(f func(s string) string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// Timeout sets the maximum duration an Action may take to invoke.
//
//	carapace.ActionCallback(func(c carapace.Context) carapace.Action {
//		time.Sleep(2*time.Second)
//		return carapace.ActionValues("done")
//	}).Timeout(1*time.Second, carapace.ActionMessage("timeout exceeded"))
func (a Action) Timeout(d time.Duration, alternative Action) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// Unique ensures the Action only contains unique values.
func (a Action) Unique() Action { _ = "STUB: not implemented"; return *new(Action) }

// UniqueList wraps the Action in an ActionMultiParts with given divider.
func (a Action) UniqueList(divider string) Action { _ = "STUB: not implemented"; return *new(Action) }

// UniqueListF is like UniqueList but uses a function to transform values before filtering.
func (a Action) UniqueListF(divider string, f func(s string) string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// Unless skips invocation if given condition is true.
func (a Action) Unless(condition bool) Action { _ = "STUB: not implemented"; return *new(Action) }

// UnlessF skips invocation if given condition returns true.
func (a Action) UnlessF(condition func(c Context) bool) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// Uid TODO experimental
func (a Action) Uid(scheme, host string, opts ...string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// UidF TODO experimental
func (a Action) UidF(f func(s string, uc uid.Context) (*url.URL, error)) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// Usage sets the usage.
func (a Action) Usage(usage string, args ...any) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// avoids failed inferred printf check: non-constant format string

// Usage sets the usage using a function.
func (a Action) UsageF(f func() string) Action { _ = "STUB: not implemented"; return *new(Action) }

// Query TODO experimental
func (a Action) Query(scheme, host, path string, opts ...string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// implicitly skip empty values

// QueryF TODO experimental
func (a Action) QueryF(f func(s string, uc uid.Context) (*url.URL, error)) Action {
	_ = "STUB: not implemented" // TODO remove the string
	return *new(Action)
}

// TODO string parameter from uid isn't really needed here

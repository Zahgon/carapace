package bash

// RedirectError current position is a redirect like `echo test >[TAB]`.
type RedirectError struct{}

func (r RedirectError) Error() string { _ = "STUB: not implemented"; return "" }

// TODO yuck! - set by Patch which also unsets bash comp environment variables so that they don't affect further completion
// introduces state and hides what is happening but works for now
var wordbreakPrefix string = ""
var compType = ""

const (
	COMP_TYPE_NORMAL               = "9"  // TAB, for normal completion
	COMP_TYPE_LIST_PARTIAL_WORD    = "33" // ‘!’, for listing alternatives on partial word completion,
	COMP_TYPE_MENU_COMPLETION      = "37" // ‘%’, for menu completion
	COMP_TYPE_LIST_SUCCESSIVE_TABS = "63" // ‘?’, for listing completions after successive tabs,
	COMP_TYPE_LIST_NOT_UNMODIFIED  = "64" // ‘@’, to list completions if the word is not unmodified
)

func CompLine() (string, bool) { _ = "STUB: not implemented"; return "", false }

// Patch patches args if `COMP_LINE` environment variable is set.
//
// Bash passes redirects to the completion function so these need to be filtered out.
//
//	`example action >/tmp/stdout.txt --values 2>/tmp/stderr.txt fi[TAB]`
//	["example", "action", ">", "/tmp/stdout.txt", "--values", "2", ">", "/tmp/stderr.txt", "fi"]
//	["example", "action", "--values", "fi"]
func Patch(args []string) ([]string, error) {
	_ = "STUB: not implemented" // TODO document and fix wordbreak splitting (e.g. `:`)
	return nil, nil
}

// TODO find a better solution to pass the wordbreakprefix to bash/action.go

func unsetBashCompEnv() { _ = "STUB: not implemented"; return }

// https://www.gnu.org/software/bash/manual/html_node/Bash-Variables.html

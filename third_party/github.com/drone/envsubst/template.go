package envsubst

import (
	"io"

	"github.com/carapace-sh/carapace/third_party/github.com/drone/envsubst/parse"
)

// state represents the state of template execution. It is not part of the
// template so that multiple executions can run in parallel.
type state struct {
	template *Template
	writer   io.Writer
	node     parse.Node // current node

	// maps variable names to values
	mapper func(string) string
}

// Template is the representation of a parsed shell format string.
type Template struct {
	tree *parse.Tree
}

// Parse creates a new shell format template and parses the template
// definition from string s.
func Parse(s string) (t *Template, err error) { _ = "STUB: not implemented"; return nil, nil }

// ParseFile creates a new shell format template and parses the template
// definition from the named file.
func ParseFile(path string) (*Template, error) { _ = "STUB: not implemented"; return nil, nil }

// Execute applies a parsed template to the specified data mapping.
func (t *Template) Execute(mapping func(string) string) (str string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *Template) eval(s *state) (err error) { _ = "STUB: not implemented"; return nil }

func (t *Template) evalText(s *state, node *parse.TextNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Template) evalList(s *state, node *parse.ListNode) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *Template) evalFunc(s *state, node *parse.FuncNode) error {
	_ = "STUB: not implemented"
	return nil
}

// restore the origin writer

// lookupFunc returns the parameters substitution function by name. If the
// named function does not exists, a default function is returned.
func lookupFunc(name string, args int) substituteFunc {
	_ = "STUB: not implemented"
	return *new(substituteFunc)
}

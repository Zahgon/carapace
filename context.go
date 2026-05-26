package carapace

import (
	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/spf13/cobra"
)

// Context provides information during completion.
type Context struct {
	// Value contains the value currently being completed (or part of it during an ActionMultiParts).
	Value string
	// Args contains the positional arguments of current (sub)command (exclusive the one currently being completed).
	Args []string
	// Parts contains the splitted Value during an ActionMultiParts (exclusive the part currently being completed).
	Parts []string
	// Env contains environment variables for current context.
	Env []string
	// Dir contains the working directory for current context.
	Dir string

	mockedReplies map[string]string
	cmd           *cobra.Command // needed for ActionCobra
}

// NewContext creates a new context for given arguments.
func NewContext(args ...string) Context { _ = "STUB: not implemented"; return *new(Context) }

// LookupEnv retrieves the value of the environment variable named by the key.
func (c Context) LookupEnv(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// Getenv retrieves the value of the environment variable named by the key.
func (c Context) Getenv(key string) string { _ = "STUB: not implemented"; return "" }

// Setenv sets the value of the environment variable named by the key.
func (c *Context) Setenv(key, value string) { _ = "STUB: not implemented"; return }

// Envsubst replaces ${var} in the string based on environment variables in current context.
func (c Context) Envsubst(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Command returns the Cmd struct to execute the named program with the given arguments.
// Env and Dir are set using the Context.
// See exec.Command for most details.
func (c Context) Command(name string, arg ...string) *execlog.Cmd {
	_ = "STUB: not implemented"
	return nil
}

// TODO use mock

func expandHome(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Abs returns an absolute representation of path.
func (c Context) Abs(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// path is relative

// prevent `C:` -> `C:./current/working/directory`

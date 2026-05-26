package execlog

import (
	"github.com/carapace-sh/carapace/third_party/golang.org/x/sys/execabs"
)

type Cmd struct {
	*execabs.Cmd
}

// Command is like execabs.Command but logs args on execution.
func Command(name string, arg ...string) *Cmd { _ = "STUB: not implemented"; return nil }

func (c *Cmd) CombinedOutput() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Cmd) Output() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Cmd) Run() error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Start() error { _ = "STUB: not implemented"; return nil }

// Command is the same as execabs.Command.
func LookPath(file string) (string, error) { _ = "STUB: not implemented"; return "", nil }

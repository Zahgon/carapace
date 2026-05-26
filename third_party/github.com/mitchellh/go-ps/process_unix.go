//go:build linux || solaris

package ps

// UnixProcess is an implementation of Process that contains Unix-specific
// fields and information.
type UnixProcess struct {
	pid   int
	ppid  int
	state rune
	pgrp  int
	sid   int

	binary string
}

func (p *UnixProcess) Pid() int { _ = "STUB: not implemented"; return 0 }

func (p *UnixProcess) PPid() int { _ = "STUB: not implemented"; return 0 }

func (p *UnixProcess) Executable() string { _ = "STUB: not implemented"; return "" }

func findProcess(pid int) (Process, error) { _ = "STUB: not implemented"; return *new(Process), nil }

func processes() ([]Process, error) { _ = "STUB: not implemented"; return nil, nil }

// We only care if the name starts with a numeric

// From this point forward, any errors we just ignore, because
// it might simply be that the process doesn't exist anymore.

func newUnixProcess(pid int) (*UnixProcess, error) { _ = "STUB: not implemented"; return nil, nil }

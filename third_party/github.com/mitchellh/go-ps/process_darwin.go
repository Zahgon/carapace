//go:build darwin

package ps

import (
	"bytes"
)

type DarwinProcess struct {
	pid    int
	ppid   int
	binary string
}

func (p *DarwinProcess) Pid() int { _ = "STUB: not implemented"; return 0 }

func (p *DarwinProcess) PPid() int { _ = "STUB: not implemented"; return 0 }

func (p *DarwinProcess) Executable() string { _ = "STUB: not implemented"; return "" }

func findProcess(pid int) (Process, error) { _ = "STUB: not implemented"; return *new(Process), nil }

func processes() ([]Process, error) { _ = "STUB: not implemented"; return nil, nil }

func darwinCstring(s [16]byte) string { _ = "STUB: not implemented"; return "" }

func darwinSyscall() (*bytes.Buffer, error) { _ = "STUB: not implemented"; return nil, nil }

const (
	_CTRL_KERN         = 1
	_KERN_PROC         = 14
	_KERN_PROC_ALL     = 0
	_KINFO_STRUCT_SIZE = 648
)

type kinfoProc struct {
	_    [40]byte
	Pid  int32
	_    [199]byte
	Comm [16]byte
	_    [301]byte
	PPid int32
	_    [84]byte
}

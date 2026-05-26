package carapace

import (
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// TODO storage needs better naming and structure

type entry struct {
	flag          ActionMap
	flagMutex     sync.RWMutex
	positional    []Action
	positionalAny *Action
	dash          []Action
	dashAny       *Action
	preinvoke     func(cmd *cobra.Command, flag *pflag.Flag, action Action) Action
	prerun        func(cmd *cobra.Command, args []string)
	bridged       bool
	initialized   bool
}

type _storage map[*cobra.Command]*entry

var storageMutex sync.RWMutex

func (s _storage) get(cmd *cobra.Command) *entry { _ = "STUB: not implemented"; return nil }

var bridgeMutex sync.Mutex

func (s _storage) bridge(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (s _storage) hasFlag(cmd *cobra.Command, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s _storage) getFlag(cmd *cobra.Command, name string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// TODO verify order of execution is correct

func (s _storage) preRun(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func (s _storage) preinvoke(cmd *cobra.Command, flag *pflag.Flag, action Action) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

func (s _storage) hasPositional(cmd *cobra.Command, index int) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO fallback to cobra defined completion if exists

func (s _storage) getPositional(cmd *cobra.Command, index int) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

func (s _storage) check() []string { _ = "STUB: not implemented"; return nil }

var storage = make(_storage)

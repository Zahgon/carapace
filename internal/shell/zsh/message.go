package zsh

import (
	"github.com/carapace-sh/carapace/internal/common"
)

type message struct {
	common.Meta
}

func (m message) Format() string { _ = "STUB: not implemented"; return "" }

func (m message) formatMessage(message, _style string) string { _ = "STUB: not implemented"; return "" }

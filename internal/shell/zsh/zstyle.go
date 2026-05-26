package zsh

import (
	"github.com/carapace-sh/carapace/internal/common"
)

type zstyles struct {
	rawValues common.RawValues
}

func (z zstyles) descriptionSGR() string { _ = "STUB: not implemented"; return "" }

func (z zstyles) valueSGR(val common.RawValue) string { _ = "STUB: not implemented"; return "" }

func (z zstyles) Format() string { _ = "STUB: not implemented"; return "" }

// disable styling for large amount of values (bad performance)

// match value with description

// only match value (also matches aliased completions that are placed on the same line if the space allows it)

// match description for aliased completions

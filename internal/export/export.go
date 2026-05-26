package export

import (
	"github.com/carapace-sh/carapace/internal/common"
)

type Export struct {
	Version string `json:"version"`
	common.Meta
	Values common.RawValues `json:"values"`
}

func (e Export) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func version() string { _ = "STUB: not implemented"; return "" }

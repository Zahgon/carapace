package nushell

type nushellStyle struct {
	Foreground string `json:"fg,omitempty"`
	Background string `json:"bg,omitempty"`
	Attributes string `json:"attr,omitempty"`
}

func convertStyle(_style string) *nushellStyle { _ = "STUB: not implemented"; return nil }

func convertColor(color string) string { _ = "STUB: not implemented"; return "" }

// keep hex

//  TODO should differfrom bright-white

// TODO

// TODO should differ from white

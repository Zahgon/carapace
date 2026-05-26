package ui

var sgrStyling = map[int]Styling{
	0: Reset,
	1: Bold,
	2: Dim,
	4: Underlined,
	5: Blink,
	7: Inverse,
}

// StyleFromSGR builds a Style from an SGR sequence.
func StyleFromSGR(s string) Style { _ = "STUB: not implemented"; return *new(Style) }

// StylingFromSGR builds a Style from an SGR sequence.
func StylingFromSGR(s string) Styling { _ = "STUB: not implemented"; return *new(Styling) }

// Do nothing; skip this code

func getSGRCodes(s string) []int { _ = "STUB: not implemented"; return nil }

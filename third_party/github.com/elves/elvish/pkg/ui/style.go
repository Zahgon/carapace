package ui

// Style specifies how something (mostly a string) shall be displayed.
type Style struct {
	Foreground Color
	Background Color
	Bold       bool
	Dim        bool
	Italic     bool
	Underlined bool
	Blink      bool
	Inverse    bool
}

// SGR returns SGR sequence for the style.
func (s Style) SGR() string { _ = "STUB: not implemented"; return "" }

// MergeFromOptions merges all recognized values from a map to the current
// Style.
func (s *Style) MergeFromOptions(options map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

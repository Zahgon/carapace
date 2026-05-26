package ui

// Styling specifies how to change a Style. It can also be applied to a Segment
// or Text.
type Styling interface{ transform(*Style) }

// ApplyStyling returns a new Style with the given Styling's applied.
func ApplyStyling(s Style, ts ...Styling) Style { _ = "STUB: not implemented"; return *new(Style) }

// Stylings joins several transformers into one.
func Stylings(ts ...Styling) Styling {
	_ = "STUB: not implemented"
	return *

	// Common stylings.
	new(Styling)
}

var (
	Reset Styling = reset{}

	FgDefault Styling = setForeground{nil}

	FgBlack   Styling = setForeground{Black}
	FgRed     Styling = setForeground{Red}
	FgGreen   Styling = setForeground{Green}
	FgYellow  Styling = setForeground{Yellow}
	FgBlue    Styling = setForeground{Blue}
	FgMagenta Styling = setForeground{Magenta}
	FgCyan    Styling = setForeground{Cyan}
	FgWhite   Styling = setForeground{White}

	FgBrightBlack   Styling = setForeground{BrightBlack}
	FgBrightRed     Styling = setForeground{BrightRed}
	FgBrightGreen   Styling = setForeground{BrightGreen}
	FgBrightYellow  Styling = setForeground{BrightYellow}
	FgBrightBlue    Styling = setForeground{BrightBlue}
	FgBrightMagenta Styling = setForeground{BrightMagenta}
	FgBrightCyan    Styling = setForeground{BrightCyan}
	FgBrightWhite   Styling = setForeground{BrightWhite}

	BgDefault Styling = setBackground{nil}

	BgBlack   Styling = setBackground{Black}
	BgRed     Styling = setBackground{Red}
	BgGreen   Styling = setBackground{Green}
	BgYellow  Styling = setBackground{Yellow}
	BgBlue    Styling = setBackground{Blue}
	BgMagenta Styling = setBackground{Magenta}
	BgCyan    Styling = setBackground{Cyan}
	BgWhite   Styling = setBackground{White}

	BgBrightBlack   Styling = setBackground{BrightBlack}
	BgBrightRed     Styling = setBackground{BrightRed}
	BgBrightGreen   Styling = setBackground{BrightGreen}
	BgBrightYellow  Styling = setBackground{BrightYellow}
	BgBrightBlue    Styling = setBackground{BrightBlue}
	BgBrightMagenta Styling = setBackground{BrightMagenta}
	BgBrightCyan    Styling = setBackground{BrightCyan}
	BgBrightWhite   Styling = setBackground{BrightWhite}

	Bold       Styling = boolOn{boldField{}}
	Dim        Styling = boolOn{dimField{}}
	Italic     Styling = boolOn{italicField{}}
	Underlined Styling = boolOn{underlinedField{}}
	Blink      Styling = boolOn{blinkField{}}
	Inverse    Styling = boolOn{inverseField{}}

	NoBold       Styling = boolOff{boldField{}}
	NoDim        Styling = boolOff{dimField{}}
	NoItalic     Styling = boolOff{italicField{}}
	NoUnderlined Styling = boolOff{underlinedField{}}
	NoBlink      Styling = boolOff{blinkField{}}
	NoInverse    Styling = boolOff{inverseField{}}

	ToggleBold       Styling = boolToggle{boldField{}}
	ToggleDim        Styling = boolToggle{dimField{}}
	ToggleItalic     Styling = boolToggle{italicField{}}
	ToggleUnderlined Styling = boolToggle{underlinedField{}}
	ToggleBlink      Styling = boolToggle{blinkField{}}
	ToggleInverse    Styling = boolToggle{inverseField{}}
)

// Fg returns a Styling that sets the foreground color.
func Fg(c Color) Styling {
	_ = "STUB: not implemented"
	return *

	// Bg returns a Styling that sets the background color.
	new(Styling)
}

func Bg(c Color) Styling { _ = "STUB: not implemented"; return *new(Styling) }

type reset struct{}
type setForeground struct{ c Color }
type setBackground struct{ c Color }
type boolOn struct{ f boolField }
type boolOff struct{ f boolField }
type boolToggle struct{ f boolField }

func (reset) transform(s *Style)           { _ = "STUB: not implemented"; return }
func (t setForeground) transform(s *Style) { _ = "STUB: not implemented"; return }
func (t setBackground) transform(s *Style) { _ = "STUB: not implemented"; return }
func (t boolOn) transform(s *Style)        { _ = "STUB: not implemented"; return }
func (t boolOff) transform(s *Style)       { _ = "STUB: not implemented"; return }
func (t boolToggle) transform(s *Style)    { _ = "STUB: not implemented"; return }

type boolField interface{ get(*Style) *bool }

type boldField struct{}
type dimField struct{}
type italicField struct{}
type underlinedField struct{}
type blinkField struct{}
type inverseField struct{}

func (boldField) get(s *Style) *bool       { _ = "STUB: not implemented"; return nil }
func (dimField) get(s *Style) *bool        { _ = "STUB: not implemented"; return nil }
func (italicField) get(s *Style) *bool     { _ = "STUB: not implemented"; return nil }
func (underlinedField) get(s *Style) *bool { _ = "STUB: not implemented"; return nil }
func (blinkField) get(s *Style) *bool      { _ = "STUB: not implemented"; return nil }
func (inverseField) get(s *Style) *bool    { _ = "STUB: not implemented"; return nil }

type jointStyling []Styling

func (t jointStyling) transform(s *Style) { _ = "STUB: not implemented"; return }

// ParseStyling parses a text representation of Styling, which are kebab
// case counterparts to the names of the builtin Styling's. For example,
// ToggleInverse is expressed as "toggle-inverse".
//
// Multiple stylings can be joined by spaces, which is equivalent to calling
// Stylings.
//
// If the given string is invalid, ParseStyling returns nil.
func ParseStyling(s string) Styling { _ = "STUB: not implemented"; return *new(Styling) }

var boolFields = map[string]boolField{
	"bold":       boldField{},
	"dim":        dimField{},
	"italic":     italicField{},
	"underlined": underlinedField{},
	"blink":      blinkField{},
	"inverse":    inverseField{},
}

func parseOneStyling(name string) Styling { _ = "STUB: not implemented"; return *new(Styling) }

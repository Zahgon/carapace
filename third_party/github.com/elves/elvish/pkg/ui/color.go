package ui

// Color represents a color.
type Color interface {
	fgSGR() string
	bgSGR() string
	String() string
}

// Builtin ANSI colors.
var (
	Black   Color = ansiColor(0)
	Red     Color = ansiColor(1)
	Green   Color = ansiColor(2)
	Yellow  Color = ansiColor(3)
	Blue    Color = ansiColor(4)
	Magenta Color = ansiColor(5)
	Cyan    Color = ansiColor(6)
	White   Color = ansiColor(7)

	BrightBlack   Color = ansiBrightColor(0)
	BrightRed     Color = ansiBrightColor(1)
	BrightGreen   Color = ansiBrightColor(2)
	BrightYellow  Color = ansiBrightColor(3)
	BrightBlue    Color = ansiBrightColor(4)
	BrightMagenta Color = ansiBrightColor(5)
	BrightCyan    Color = ansiBrightColor(6)
	BrightWhite   Color = ansiBrightColor(7)
)

// XTerm256Color returns a color from the xterm 256-color palette.
func XTerm256Color(i uint8) Color {
	_ = "STUB: not implemented"
	return *

	// TrueColor returns a 24-bit true color.
	new(Color)
}

func TrueColor(r, g, b uint8) Color { _ = "STUB: not implemented"; return *new(Color) }

var colorNames = []string{
	"black", "red", "green", "yellow",
	"blue", "magenta", "cyan", "white",
}

var colorByName = map[string]Color{
	"black":   Black,
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"white":   White,

	"bright-black":   BrightBlack,
	"bright-red":     BrightRed,
	"bright-green":   BrightGreen,
	"bright-yellow":  BrightYellow,
	"bright-blue":    BrightBlue,
	"bright-magenta": BrightMagenta,
	"bright-cyan":    BrightCyan,
	"bright-white":   BrightWhite,
}

type ansiColor uint8

func (c ansiColor) fgSGR() string  { _ = "STUB: not implemented"; return "" }
func (c ansiColor) bgSGR() string  { _ = "STUB: not implemented"; return "" }
func (c ansiColor) String() string { _ = "STUB: not implemented"; return "" }

type ansiBrightColor uint8

func (c ansiBrightColor) fgSGR() string  { _ = "STUB: not implemented"; return "" }
func (c ansiBrightColor) bgSGR() string  { _ = "STUB: not implemented"; return "" }
func (c ansiBrightColor) String() string { _ = "STUB: not implemented"; return "" }

type xterm256Color uint8

func (c xterm256Color) fgSGR() string  { _ = "STUB: not implemented"; return "" }
func (c xterm256Color) bgSGR() string  { _ = "STUB: not implemented"; return "" }
func (c xterm256Color) String() string { _ = "STUB: not implemented"; return "" }

type trueColor struct{ R, G, B uint8 }

func (c trueColor) fgSGR() string { _ = "STUB: not implemented"; return "" }
func (c trueColor) bgSGR() string { _ = "STUB: not implemented"; return "" }

func (c trueColor) String() string { _ = "STUB: not implemented"; return "" }

func (c trueColor) rgbSGR() string { _ = "STUB: not implemented"; return "" }

func parseColor(name string) Color { _ = "STUB: not implemented"; return *new(Color) }

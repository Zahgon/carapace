// Package style provide display coloring
package style

import (
	"github.com/carapace-sh/carapace/third_party/github.com/elves/elvish/pkg/ui"
)

var (
	Default = ""

	Black   = "black"
	Red     = "red"
	Green   = "green"
	Yellow  = "yellow"
	Blue    = "blue"
	Magenta = "magenta"
	Cyan    = "cyan"
	White   = "white"

	BrightBlack   = "bright-black"
	BrightRed     = "bright-red"
	BrightGreen   = "bright-green"
	BrightYellow  = "bright-yellow"
	BrightBlue    = "bright-blue"
	BrightMagenta = "bright-magenta"
	BrightCyan    = "bright-cyan"
	BrightWhite   = "bright-white"

	BgBlack   = "bg-black"
	BgRed     = "bg-red"
	BgGreen   = "bg-green"
	BgYellow  = "bg-yellow"
	BgBlue    = "bg-blue"
	BgMagenta = "bg-magenta"
	BgCyan    = "bg-cyan"
	BgWhite   = "bg-white"

	BgBrightBlack   = "bg-bright-black"
	BgBrightRed     = "bg-bright-red"
	BgBrightGreen   = "bg-bright-green"
	BgBrightYellow  = "bg-bright-yellow"
	BgBrightBlue    = "bg-bright-blue"
	BgBrightMagenta = "bg-bright-magenta"
	BgBrightCyan    = "bg-bright-cyan"
	BgBrightWhite   = "bg-bright-white"

	Bold       = "bold"
	Dim        = "dim"
	Italic     = "italic"
	Underlined = "underlined"
	Blink      = "blink"
	Inverse    = "inverse"
)

// Of combines different styles.
func Of(s ...string) string { _ = "STUB: not implemented"; return "" }

// XTerm256Color returns a color from the xterm 256-color palette.
func XTerm256Color(i uint8) string { _ = "STUB: not implemented"; return "" }

// TrueColor returns a 24-bit true color.
func TrueColor(r, g, b uint8) string { _ = "STUB: not implemented"; return "" }

// SGR returns the SGR sequence for given style.
func SGR(s string) string { _ = "STUB: not implemented"; return "" }

func Parse(s string) ui.Style { _ = "STUB: not implemented"; return *new(ui.Style) }

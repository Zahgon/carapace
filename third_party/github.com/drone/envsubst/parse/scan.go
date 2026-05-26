package parse

// eof rune sent when end of file is reached
var eof = rune(0)

// token is a lexical token.
type token uint

// list of lexical tokens.
const (
	// special tokens
	tokenIllegal token = iota
	tokenEOF

	// identifiers and literals
	tokenIdent

	// operators and delimiters
	tokenLbrack
	tokenRbrack
	tokenQuote
)

// predefined mode bits to control recognition of tokens.
const (
	scanIdent byte = 1 << iota
	scanLbrack
	scanRbrack
	scanEscape
)

// predefined mode bits to control escape tokens.
const (
	dollar byte = 1 << iota
	backslash
	escapeAll = dollar | backslash
)

// returns true if rune is accepted.
type acceptFunc func(r rune, i int) bool

// scanner implements a lexical scanner that reads unicode
// characters and tokens from a string buffer.
type scanner struct {
	buf         string
	pos         int
	start       int
	width       int
	mode        byte
	escapeChars byte

	accept acceptFunc
}

// init initializes a scanner with a new buffer.
func (s *scanner) init(buf string) {
	s.buf = buf
	s.pos = 0
	s.start = 0
	s.width = 0
	s.accept = nil
}

// read returns the next unicode character. It returns eof at
// the end of the string buffer.
func (s *scanner) read() rune { _ = "STUB: not implemented"; return 0 }

func (s *scanner) unread() {
	_ = "STUB: not implemented"

	// skip skips over the curring unicode character in the buffer
	// by slicing and removing from the buffer.
	return
}

func (s *scanner) skip() { _ = "STUB: not implemented"; return }

// peek returns the next unicode character in the buffer without
// advancing the scanner. It returns eof if the scanner's position
// is at the last character of the source.
func (s *scanner) peek() rune { _ = "STUB: not implemented"; return 0 }

// string returns the string corresponding to the most recently
// scanned token. Valid after calling scan().
func (s *scanner) string() string { _ = "STUB: not implemented"; return "" }

// tests if the bit exists for a given character bit
func (s *scanner) shouldEscape(character byte) bool { _ = "STUB: not implemented"; return false }

// scan reads the next token or Unicode character from source and
// returns it. It returns EOF at the end of the source.
func (s *scanner) scan() token { _ = "STUB: not implemented"; return *new(token) }

// scanIdent reads the next token or Unicode character from source
// and returns true if the Ident character is accepted.
func (s *scanner) scanIdent(r rune) bool { _ = "STUB: not implemented"; return false }

// scanLbrack reads the next token or Unicode character from source
// and returns true if the open bracket is encountered.
func (s *scanner) scanLbrack(r rune) bool { _ = "STUB: not implemented"; return false }

// scanRbrack reads the next token or Unicode character from source
// and returns true if the closing bracket is encountered.
func (s *scanner) scanRbrack(r rune) bool { _ = "STUB: not implemented"; return false }

// scanEscaped reads the next token or Unicode character from source
// and returns true if it being escaped and should be skipped.
func (s *scanner) scanEscaped(r rune) bool { _ = "STUB: not implemented"; return false }

//
// scanner functions accept or reject runes.
//

func acceptRune(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptIdent(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptColon(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptOneExclamationMark(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptOneHash(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptNotClosing(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptHashFunc(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptPercentFunc(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptDefaultFunc(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptReplaceFunc(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptOneEqual(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptOneColon(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func rejectColonClose(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptSlash(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptNotSlash(r rune, i int) bool { _ = "STUB: not implemented"; return false }

func acceptCasingFunc(r rune, i int) bool { _ = "STUB: not implemented"; return false }

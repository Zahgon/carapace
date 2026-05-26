package jsonc

import (
	"regexp"
)

// jsonc is the structure for parsing json with comments
type jsonc struct {
	index    int    // current index position in source data
	comment  uint   // the type of comment: eithher 1 for // OR 2 for /*
	len      int    // the length of source data
	objDepth uint   // the depth of nested objects
	arrDepth uint   // the depth of nested arrays
	inStr    bool   // if inside string
	inArr    bool   // if inside array notation: [
	inObj    bool   // if inside object notation: {
	last     string // last significant non whitespace char
	strDelim string // string delimeter: either ' or "
}

// new creates Jsonc struct
func new() *jsonc {
	_ = "STUB: not implemented"

	// Strip strips comments and trailing commas from input byte array
	return nil
}

func Strip(jsonb []byte) []byte { _ = "STUB: not implemented"; return nil }

var sq = `'`  // single quote
var dq = `"`  // double quote
var esc = `\` // escape
var comma = regexp.MustCompile(`(?:,+)(\s*)$`)

// StripS strips comments and trailing commas from input string
func (j *jsonc) StripS(data string) string { _ = "STUB: not implemented"; return "" }

// If value starts with 0x, parse as hexadecimal

// Trim trailing commas at the end of array or object

// Append char as is (or it's compliment pair) if inside string or outside comment

// Wipe out trailing whitespaces around comment

// Unmarshal strips and parses the json byte array
func Unmarshal(jsonb []byte, v any) error { _ = "STUB: not implemented"; return nil }

// reset resets the Jsonc with proper defaults
func (j *jsonc) reset() { _ = "STUB: not implemented"; return }

// getSegments gets look-behind, current and look-ahead chars
func (j *jsonc) getSegments(json, old string) (oldprev, prev, char, next string) {
	_ = "STUB: not implemented"
	return "", "", "", ""
}

// isNonStringValue checks if char is value outside string (or comment) and matches chars
func (j *jsonc) isNonStringValue(char, chars string) bool { _ = "STUB: not implemented"; return false }

// hexadecimal consumes hex (0-9a-fA-F) chars and converts to decimal string
func (j *jsonc) hexadecimal(data string) string { _ = "STUB: not implemented"; return "" }

// quoteKey double quotes the unquoted object keys
func (j *jsonc) quoteKey(char string, wasQuoted bool) (q string, quoted bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Object key has just started without quote, so quote it

// Object key has just ended and was quoted before, so quote it again to compliment

// checkArrayObject checks and sets the depth and state of array &/or object notation
func (j *jsonc) checkArrayObject(char string) { _ = "STUB: not implemented"; return }

// Last non whitespace char

func (j *jsonc) inString(prev, char, next, oldprev string) bool {
	_ = "STUB: not implemented"
	return false
}

// Toggle j.inStr if j.strDelim is not escaped

// outsideComment checks if char is outside comment
// it also sets the state of comment
func (j *jsonc) outsideComment(char, next string) bool {
	_ = "STUB: not implemented"
	// Set comment state: `//` => 1 | `/*` => 2
	return false
}

// hasCommentEnded checks if the comment has just ended and resets the state
func (j *jsonc) hasCommentEnded(char, next string) bool {
	_ = "STUB: not implemented"
	// Single line comment ends with `\n` and multiline ends with `*/`
	return false
}

var spacesPair = map[string]string{"\n": `\n`, "\t": `\t`, "\r": `\r`}

// compliment appends char as is (or it's compliment pair)
// (eg: in string boundary the compliment of single quote is double quote)
// it also normalizes whitespaces inside string and signed &/or decimal numbers
func (j *jsonc) compliment(prev, char, next string) string { _ = "STUB: not implemented"; return "" }

// Signed +ve number

// Decimal point number

// Single quoted string

// isNumber checks if a string char is numeric
func isNumber(char string, hex bool) bool { _ = "STUB: not implemented"; return false }

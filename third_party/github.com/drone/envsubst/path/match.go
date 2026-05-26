// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package path

import (
	"errors"
)

// ErrBadPattern indicates a globbing pattern was malformed.
var ErrBadPattern = errors.New("syntax error in pattern")

// Match reports whether name matches the shell file name pattern.
// The pattern syntax is:
//
//	pattern:
//		{ term }
//	term:
//		'*'         matches any sequence of non-/ characters
//		'?'         matches any single non-/ character
//		'[' [ '^' ] { character-range } ']'
//		            character class (must be non-empty)
//		c           matches character c (c != '*', '?', '\\', '[')
//		'\\' c      matches character c
//
//	character-range:
//		c           matches character c (c != '\\', '-', ']')
//		'\\' c      matches character c
//		lo '-' hi   matches character c for lo <= c <= hi
//
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is ErrBadPattern, when pattern
// is malformed.
func Match(pattern, name string) (matched bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Trailing * matches rest of string unless it has a /.
// return !strings.Contains(name, "/"), nil

// Return rest of string

// Look for match at current position.

// if we're the last chunk, make sure we've exhausted the name
// otherwise we'll give a false result even if we could still match
// using the star

// Look for match skipping i+1 bytes.

// if we're the last chunk, make sure we exhausted the name

// scanChunk gets the next segment of pattern, which is a non-star string
// possibly preceded by a star.
func scanChunk(pattern string) (star bool, chunk, rest string) {
	_ = "STUB: not implemented"
	return false, "", ""
}

// error check handled in matchChunk: bad pattern.

// matchChunk checks whether chunk matches the beginning of s.
// If so, it returns the remainder of s (after the match).
// Chunk is all single-character operators: literals, char classes, and ?.
func matchChunk(chunk, s string) (rest string, ok bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// character class

// possibly negated

// parse all ranges

// getEsc gets a possibly-escaped character from chunk, for a character class.
func getEsc(chunk string) (r rune, nchunk string, err error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

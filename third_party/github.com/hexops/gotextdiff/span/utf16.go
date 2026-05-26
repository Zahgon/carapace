// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package span

// ToUTF16Column calculates the utf16 column expressed by the point given the
// supplied file contents.
// This is used to convert from the native (always in bytes) column
// representation and the utf16 counts used by some editors.
func ToUTF16Column(p Point, content []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// 0-based
// 0-based

// 0-based column 0, so it must be chr 1

// work out the offset at the start of the line using the column

// Use the offset to pick out the line start.
// This cannot panic: offset > len(content) and lineOffset < offset.

// Now, truncate down to the supplied column.

// and count the number of utf16 characters
// in theory we could do this by hand more efficiently...

// FromUTF16Column advances the point by the utf16 character offset given the
// supplied line contents.
// This is used to convert from the utf16 counts used by some editors to the
// native (always in bytes) column representation.
func FromUTF16Column(p Point, chr int, content []byte) (Point, error) {
	_ = "STUB: not implemented"
	return *new(Point), nil
}

// if chr is 1 then no adjustment needed

// scan forward the specified number of characters

// Per the LSP spec:
//
// > If the character value is greater than the line length it
// > defaults back to the line length.

// a two point rune

// if we finished in a two point rune, do not advance past the first

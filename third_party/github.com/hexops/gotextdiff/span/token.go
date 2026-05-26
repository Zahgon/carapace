// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package span

import (
	"go/token"
)

// Range represents a source code range in token.Pos form.
// It also carries the FileSet that produced the positions, so that it is
// self contained.
type Range struct {
	FileSet   *token.FileSet
	Start     token.Pos
	End       token.Pos
	Converter Converter
}

type FileConverter struct {
	file *token.File
}

// TokenConverter is a Converter backed by a token file set and file.
// It uses the file set methods to work out the conversions, which
// makes it fast and does not require the file contents.
type TokenConverter struct {
	FileConverter
	fset *token.FileSet
}

// NewRange creates a new Range from a FileSet and two positions.
// To represent a point pass a 0 as the end pos.
func NewRange(fset *token.FileSet, start, end token.Pos) Range {
	_ = "STUB: not implemented"
	return *new(Range)
}

// NewTokenConverter returns an implementation of Converter backed by a
// token.File.
func NewTokenConverter(fset *token.FileSet, f *token.File) *TokenConverter {
	_ = "STUB: not implemented"
	return nil
}

// NewContentConverter returns an implementation of Converter for the
// given file content.
func NewContentConverter(filename string, content []byte) *TokenConverter {
	_ = "STUB: not implemented"
	return nil
}

// IsPoint returns true if the range represents a single point.
func (r Range) IsPoint() bool { _ = "STUB: not implemented"; return false }

// Span converts a Range to a Span that represents the Range.
// It will fill in all the members of the Span, calculating the line and column
// information.
func (r Range) Span() (Span, error) { _ = "STUB: not implemented"; return *new(Span), nil }

// FileSpan returns a span within tok, using converter to translate between
// offsets and positions.
func FileSpan(tok *token.File, converter Converter, start, end token.Pos) (Span, error) {
	_ = "STUB: not implemented"
	return *new(Span), nil
}

// In the presence of line directives, a single File can have sections from
// multiple file names.

func position(f *token.File, pos token.Pos) (string, int, int, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

func positionFromOffset(f *token.File, offset int) (string, int, int, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

// TODO(golang/go#41029): Consider returning line, column instead of line+1, 1 if
// the file's last character is not a newline.

// offset is a copy of the Offset function in go/token, but with the adjustment
// that it does not panic on invalid positions.
func offset(f *token.File, pos token.Pos) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Range converts a Span to a Range that represents the Span for the supplied
// File.
func (s Span) Range(converter *TokenConverter) (Range, error) {
	_ = "STUB: not implemented"
	return *new(Range), nil
}

// go/token will panic if the offset is larger than the file's size,
// so check here to avoid panicking.

func (l *FileConverter) ToPosition(offset int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (l *FileConverter) ToOffset(line, col int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// at the end of the file, allowing for a trailing eol

// we assume that column is in bytes here, and that the first byte of a
// line is at column 1

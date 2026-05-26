// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package span contains support for representing with positions and ranges in
// text files.
package span

import (
	"fmt"
)

// Span represents a source code range in standardized form.
type Span struct {
	v span
}

// Point represents a single point within a file.
// In general this should only be used as part of a Span, as on its own it
// does not carry enough information.
type Point struct {
	v point
}

type span struct {
	URI   URI   `json:"uri"`
	Start point `json:"start"`
	End   point `json:"end"`
}

type point struct {
	Line   int `json:"line"`
	Column int `json:"column"`
	Offset int `json:"offset"`
}

// Invalid is a span that reports false from IsValid
var Invalid = Span{v: span{Start: invalidPoint.v, End: invalidPoint.v}}

var invalidPoint = Point{v: point{Line: 0, Column: 0, Offset: -1}}

// Converter is the interface to an object that can convert between line:column
// and offset forms for a single file.
type Converter interface {
	//ToPosition converts from an offset to a line:column pair.
	ToPosition(offset int) (int, int, error)
	//ToOffset converts from a line:column pair to an offset.
	ToOffset(line, col int) (int, error)
}

func New(uri URI, start Point, end Point) Span { _ = "STUB: not implemented"; return *new(Span) }

func NewPoint(line, col, offset int) Point { _ = "STUB: not implemented"; return *new(Point) }

func Compare(a, b Span) int { _ = "STUB: not implemented"; return 0 }

func ComparePoint(a, b Point) int { _ = "STUB: not implemented"; return 0 }

func comparePoint(a, b point) int { _ = "STUB: not implemented"; return 0 }

func (s Span) HasPosition() bool             { _ = "STUB: not implemented"; return false }
func (s Span) HasOffset() bool               { _ = "STUB: not implemented"; return false }
func (s Span) IsValid() bool                 { _ = "STUB: not implemented"; return false }
func (s Span) IsPoint() bool                 { _ = "STUB: not implemented"; return false }
func (s Span) URI() URI                      { _ = "STUB: not implemented"; return *new(URI) }
func (s Span) Start() Point                  { _ = "STUB: not implemented"; return *new(Point) }
func (s Span) End() Point                    { _ = "STUB: not implemented"; return *new(Point) }
func (s *Span) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (s *Span) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (p Point) HasPosition() bool             { _ = "STUB: not implemented"; return false }
func (p Point) HasOffset() bool               { _ = "STUB: not implemented"; return false }
func (p Point) IsValid() bool                 { _ = "STUB: not implemented"; return false }
func (p *Point) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (p *Point) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
func (p Point) Line() int                     { _ = "STUB: not implemented"; return 0 }

func (p Point) Column() int { _ = "STUB: not implemented"; return 0 }

func (p Point) Offset() int { _ = "STUB: not implemented"; return 0 }

func (p point) hasPosition() bool { _ = "STUB: not implemented"; return false }
func (p point) hasOffset() bool   { _ = "STUB: not implemented"; return false }
func (p point) isValid() bool     { _ = "STUB: not implemented"; return false }
func (p point) isZero() bool      { _ = "STUB: not implemented"; return false }

func (s *span) clean() {
	_ = "STUB: not implemented"
	// this presumes the points are already clean
	return
}

func (p *point) clean() { _ = "STUB: not implemented"; return }

// Format implements fmt.Formatter to print the Location in a standard form.
// The format produced is one that can be read back in using Parse.
func (s Span) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// we should always have a uri, simplify if it is file format
//TODO: make sure the end of the uri is unambiguous

// see which bits of start to write

// start is written, do we need end?

// we don't print the line if it did not change

func (s Span) WithPosition(c Converter) (Span, error) {
	_ = "STUB: not implemented"
	return *new(Span), nil
}

func (s Span) WithOffset(c Converter) (Span, error) {
	_ = "STUB: not implemented"
	return *new(Span), nil
}

func (s Span) WithAll(c Converter) (Span, error) { _ = "STUB: not implemented"; return *new(Span), nil }

func (s *Span) update(c Converter, withPos, withOffset bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *point) updatePosition(c Converter) error { _ = "STUB: not implemented"; return nil }

func (p *point) updateOffset(c Converter) error { _ = "STUB: not implemented"; return nil }

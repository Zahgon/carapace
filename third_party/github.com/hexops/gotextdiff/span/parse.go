// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package span

// Parse returns the location represented by the input.
// Only file paths are accepted, not URIs.
// The returned span will be normalized, and thus if printed may produce a
// different string.
func Parse(input string) Span {
	_ = "STUB: not implemented"
	// :0:0#0-0:0#0
	return *new(Span)
}

// we have a span, fall out of the case to continue

// separator not valid, rewind to either the : or the start

// only the span form can get here
// at this point we still don't know what the numbers we have mean
// if have not yet seen a : then we might have either a line or a column depending
// on whether start has a column or not
// we build an end point and will fix it later if needed

// turns out we don't have a span after all, rewind

// line#offset only

// we have a column, so if end only had one number, it is also the column

type suffix struct {
	remains string
	sep     string
	num     int
}

func rstripSuffix(input string) suffix { _ = "STUB: not implemented"; return *new(suffix) }

// first see if we have a number at the end

// now see if we have a trailing separator

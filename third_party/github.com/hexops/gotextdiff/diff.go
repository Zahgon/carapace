// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package gotextdiff supports a pluggable diff algorithm.
package gotextdiff

import (
	"github.com/carapace-sh/carapace/third_party/github.com/hexops/gotextdiff/span"
)

// TextEdit represents a change to a section of a document.
// The text within the specified span should be replaced by the supplied new text.
type TextEdit struct {
	Span    span.Span
	NewText string
}

// ComputeEdits is the type for a function that produces a set of edits that
// convert from the before content to the after content.
type ComputeEdits func(uri span.URI, before, after string) []TextEdit

// SortTextEdits attempts to order all edits by their starting points.
// The sort is stable so that edits with the same starting point will not
// be reordered.
func SortTextEdits(d []TextEdit) {
	_ = "STUB: not implemented"
	// Use a stable sort to maintain the order of edits inserted at the same position.
	return
}

// ApplyEdits applies the set of edits to the before and returns the resulting
// content.
// It may panic or produce garbage if the edits are not valid for the provided
// before content.
func ApplyEdits(before string, edits []TextEdit) string {
	_ = "STUB: not implemented"
	// Preconditions:
	//   - all of the edits apply to before
	//   - and all the spans for each TextEdit have the same URI
	return ""
}

// LineEdits takes a set of edits and expands and merges them as necessary
// to ensure that there are only full line edits left when it is done.
func LineEdits(before string, edits []TextEdit) []TextEdit { _ = "STUB: not implemented"; return nil }

// prepareEdits returns a sorted copy of the edits
func prepareEdits(before string, edits []TextEdit) (*span.TokenConverter, []TextEdit, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// lineEdits rewrites the edits to always be full line edits
func lineEdits(before string, _ *span.TokenConverter, edits []TextEdit) []TextEdit {
	_ = "STUB: not implemented"
	return nil
}

// overlaps with the current edit, need to combine
// first get the gap from the previous edit

// now add the text of this edit

// and then adjust the end position

// does not overlap, add previous run (if there is one)

// and then remember this edit as the start of the next run

// add the current pending run if there is one

func addEdit(before string, edits []TextEdit, edit TextEdit) []TextEdit {
	_ = "STUB: not implemented"
	return nil
}

// if edit is partial, expand it to full line now

// prepend the text and adjust to start of line

// after end of file that does not end in eol, so join to last line of file
// to do this we need to know where the start of the last line was

// file is one non terminated line

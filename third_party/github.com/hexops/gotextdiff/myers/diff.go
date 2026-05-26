// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package myers implements the Myers diff algorithm.
package myers

import (
	diff "github.com/carapace-sh/carapace/third_party/github.com/hexops/gotextdiff"
	"github.com/carapace-sh/carapace/third_party/github.com/hexops/gotextdiff/span"
)

// Sources:
// https://blog.jcoglan.com/2017/02/17/the-myers-diff-algorithm-part-3/
// https://www.codeproject.com/Articles/42279/%2FArticles%2F42279%2FInvestigating-Myers-diff-algorithm-Part-1-of-2

func ComputeEdits(uri span.URI, before, after string) []diff.TextEdit {
	_ = "STUB: not implemented"
	return nil
}

// Delete: unformatted[i1:i2] is deleted.

// Insert: formatted[j1:j2] is inserted at unformatted[i1:i1].

type operation struct {
	Kind    diff.OpKind
	Content []string // content from b
	I1, I2  int      // indices of the line in a
	J1      int      // indices of the line in b, J2 implied by len(Content)
}

// operations returns the list of operations to convert a into b, consolidating
// operations for multiple lines and not including equal lines.
func operations(a, b []string) []*operation { _ = "STUB: not implemented"; return nil }

// delete (horizontal)

// insert (vertical)

// equal (diagonal)

// backtrack uses the trace for the edit sequence computation and returns the
// "snakes" that make up the solution. A "snake" is a single deletion or
// insertion followed by zero or diagonals.
func backtrack(trace [][]int, x, y, offset int) [][]int { _ = "STUB: not implemented"; return nil }

// shortestEditSequence returns the shortest edit sequence that converts a into b.
func shortestEditSequence(a, b []string) ([][]int, int) { _ = "STUB: not implemented"; return nil, 0 }

// Iterate through the maximum possible length of the SES (N+M).

// k lines are represented by the equation y = x - k. We move in
// increments of 2 because end points for even d are on even k lines.

// At each point, we either go down or to the right. We go down if
// k == -d, and we go to the right if k == d. We also prioritize
// the maximum x value, because we prefer deletions to insertions.

// down

// right

// Diagonal moves while we have equal contents.

// Return if we've exceeded the maximum values.

// Makes sure to save the state of the array before returning.

// Save the state of the array.

func splitLines(text string) []string { _ = "STUB: not implemented"; return nil }

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package span

const fileScheme = "file"

// URI represents the full URI for a file.
type URI string

func (uri URI) IsFile() bool { _ = "STUB: not implemented"; return false }

// Filename returns the file path for the given URI.
// It is an error to call this on a URI that is not a valid filename.
func (uri URI) Filename() string { _ = "STUB: not implemented"; return "" }

func filename(uri URI) (string, error) { _ = "STUB: not implemented"; return "", nil }

// If the URI is a Windows URI, we trim the leading "/" and lowercase
// the drive letter, which will never be case sensitive.

func URIFromURI(s string) URI { _ = "STUB: not implemented"; return *new(URI) }

// VS Code sends URLs with only two slashes, which are invalid. golang/go#39789.

// Even though the input is a URI, it may not be in canonical form. VS Code
// in particular over-escapes :, @, etc. Unescape and re-encode to canonicalize.

// File URIs from Windows may have lowercase drive letters.
// Since drive letters are guaranteed to be case insensitive,
// we change them to uppercase to remain consistent.
// For example, file:///c:/x/y/z becomes file:///C:/x/y/z.

func CompareURI(a, b URI) int { _ = "STUB: not implemented"; return 0 }

func equalURI(a, b URI) bool { _ = "STUB: not implemented"; return false }

// If we have the same URI basename, we may still have the same file URIs.

// Stat the files to check if they are equal.

// URIFromPath returns a span URI for the supplied file path.
// It will always have the file scheme.
func URIFromPath(path string) URI { _ = "STUB: not implemented"; return *new(URI) }

// Handle standard library paths that contain the literal "$GOROOT".
// TODO(rstambler): The go/packages API should allow one to determine a user's $GOROOT.

// Check the file path again, in case it became absolute.

// isWindowsDrivePath returns true if the file path is of the form used by
// Windows. We check if the path begins with a drive letter, followed by a ":".
// For example: C:/x/y/z.
func isWindowsDrivePath(path string) bool { _ = "STUB: not implemented"; return false }

// isWindowsDriveURI returns true if the file URI is of the format used by
// Windows URIs. The url.Parse package does not specially handle Windows paths
// (see golang/go#6027), so we check if the URI path has a drive prefix (e.g. "/C:").
func isWindowsDriveURIPath(uri string) bool { _ = "STUB: not implemented"; return false }

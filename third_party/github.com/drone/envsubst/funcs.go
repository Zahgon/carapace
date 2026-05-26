package envsubst

// defines a parameter substitution function.
type substituteFunc func(string, ...string) string

// toLen returns the length of string s.
func toLen(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// toLower returns a copy of the string s with all characters
// mapped to their lower case.
func toLower(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// toUpper returns a copy of the string s with all characters
// mapped to their upper case.
func toUpper(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// toLowerFirst returns a copy of the string s with the first
// character mapped to its lower case.
func toLowerFirst(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// toUpperFirst returns a copy of the string s with the first
// character mapped to its upper case.
func toUpperFirst(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// toDefault returns a copy of the string s if not empty, else
// returns a concatenation of the args without a separator.
func toDefault(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// don't use any separator

// toSubstr returns a slice of the string s at the specified
// length and position.
func toSubstr(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// should never happen

// bash returns the string if the position
// cannot be parsed.

// if pos is negative (counts from the end) add it
// to length to get first character offset

// if negative offset exceeds the length of the string
// start from 0

// if the position exceeds the length of the
// string an empty string is returned

// bash returns the string if the length
// cannot be parsed.

// if the position exceeds the length of the
// string just return the rest of it like bash

// if the position exceeds the length of the
// string an empty string is returned

// replaceAll returns a copy of the string s with all instances
// of the substring replaced with the replacement string.
func replaceAll(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// replaceFirst returns a copy of the string s with the first
// instance of the substring replaced with the replacement string.
func replaceFirst(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// replacePrefix returns a copy of the string s with the matching
// prefix replaced with the replacement string.
func replacePrefix(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// replaceSuffix returns a copy of the string s with the matching
// suffix replaced with the replacement string.
func replaceSuffix(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

// TODO

func trimShortestPrefix(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

func trimShortestSuffix(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

func trimLongestPrefix(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

func trimLongestSuffix(s string, args ...string) string { _ = "STUB: not implemented"; return "" }

func trimShortest(s, arg string) string { _ = "STUB: not implemented"; return "" }

func trimLongest(s, arg string) string { _ = "STUB: not implemented"; return "" }

func reverse(s string) string { _ = "STUB: not implemented"; return "" }

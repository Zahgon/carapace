// Package common code
package common

// FromInvokedAction provides access to RawValues within an InvokedAction.
// It is intended for testing purposes in Sandbox (circumventing dependency issues).
var FromInvokedAction func(action any) (Meta, RawValues)

// RawValue represents a completion candidate.
type RawValue struct {
	Value       string `json:"value"`
	Display     string `json:"display"`
	Description string `json:"description,omitempty"`
	Style       string `json:"style,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Uid         string `json:"uid,omitempty"`
}

// TrimmedDescription returns the trimmed description.
func (r RawValue) TrimmedDescription() string { _ = "STUB: not implemented"; return "" }

// RawValues is an alias for []RawValue.
type RawValues []RawValue

// RawValuesFrom creates RawValues from given values.
func RawValuesFrom(values ...string) RawValues { _ = "STUB: not implemented"; return *new(RawValues) }

func (r RawValues) Unique() RawValues { _ = "STUB: not implemented"; return *new(RawValues) }

func (r RawValues) contains(s string) bool { _ = "STUB: not implemented"; return false }

// Filter filters values.
func (r RawValues) Filter(values ...string) RawValues {
	_ = "STUB: not implemented"
	return *new(RawValues)
}

// Retain retains given values.
func (r RawValues) Retain(values ...string) RawValues {
	_ = "STUB: not implemented"
	return *new(RawValues)
}

// Decolor clears style for all values.
func (r RawValues) Decolor() RawValues { _ = "STUB: not implemented"; return *new(RawValues) }

// FilterPrefix filters values with given prefix.
func (r RawValues) FilterPrefix(prefix string) RawValues {
	_ = "STUB: not implemented"
	return *new(RawValues)
}

func (r RawValues) EachTag(f func(tag string, values RawValues)) { _ = "STUB: not implemented"; return }

// ByValue alias to filter by value.
type ByValue []RawValue

func (a ByValue) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a ByValue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (a ByValue) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// ByDisplay alias to filter by display.
type ByDisplay []RawValue

func (a ByDisplay) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a ByDisplay) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// ensure consistency in export

func (a ByDisplay) Swap(i, j int) { _ = "STUB: not implemented"; return }

type ByUid []RawValue

func (a ByUid) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a ByUid) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (a ByUid) Swap(i, j int)      { _ = "STUB: not implemented"; return }

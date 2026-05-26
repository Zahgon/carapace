package carapace

import (
	"net/url"

	"github.com/carapace-sh/carapace/internal/common"
	"github.com/carapace-sh/carapace/internal/export"
)

// InvokedAction is a logical alias for an Action whose (nested) callback was invoked.
type InvokedAction struct {
	action Action
}

func (ia InvokedAction) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ia InvokedAction) export() export.Export {
	_ = "STUB: not implemented"
	return *new(export.Export)
}

// Filter filters given values.
//
//	a := carapace.ActionValues("A", "B", "C").Invoke(c)
//	b := a.Filter([]string{"B"}) // ["A", "C"]
func (ia InvokedAction) Filter(values ...string) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

// Merge merges InvokedActions (existing values are overwritten)
//
//	a := carapace.ActionValues("A", "B").Invoke(c)
//	b := carapace.ActionValues("B", "C").Invoke(c)
//	c := a.Merge(b) // ["A", "B", "C"]
func (ia InvokedAction) Merge(others ...InvokedAction) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

// Prefix adds a prefix to values (only the ones inserted, not the display values)
//
//	carapace.ActionValues("melon", "drop", "fall").Invoke(c).Prefix("water")
func (ia InvokedAction) Prefix(prefix string) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

// Retain retains given values.
//
//	a := carapace.ActionValues("A", "B", "C").Invoke(c)
//	b := a.Retain([]string{"A", "C"}) // ["A", "C"]
func (ia InvokedAction) Retain(values ...string) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

// Suffix adds a suffx to values (only the ones inserted, not the display values)
//
//	carapace.ActionValues("apple", "melon", "orange").Invoke(c).Suffix("juice")
func (ia InvokedAction) Suffix(suffix string) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

// UidF TODO experimental
func (ia InvokedAction) UidF(f func(s string) (*url.URL, error)) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

// ToA casts an InvokedAction to Action.
func (ia InvokedAction) ToA() Action { _ = "STUB: not implemented"; return *new(Action) }

func tokenize(s string, dividers ...string) []string { _ = "STUB: not implemented"; return nil }

// ToMultiPartsA create an ActionMultiParts from values with given dividers
//
//	a := carapace.ActionValues("A/B/C", "A/C", "B/C", "C").Invoke(c)
//	b := a.ToMultiPartsA("/") // completes segments separately (first one is ["A/", "B/", "C"])
func (ia InvokedAction) ToMultiPartsA(dividers ...string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

func (ia InvokedAction) QueryS(q string) InvokedAction {
	_ = "STUB: not implemented"
	return *new(InvokedAction)
}

func (ia InvokedAction) value(shell string, value string) string {
	_ = "STUB: not implemented"
	return ""
}

func init() {
	common.FromInvokedAction = func(i any) (common.Meta, common.RawValues) {
		if invoked, ok := i.(InvokedAction); ok {
			return invoked.action.meta, invoked.action.rawValues
		}
		return common.Meta{}, nil
	}
}

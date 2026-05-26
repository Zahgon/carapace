package carapace

// Diff compares values of two actions.
// It overrides the style to hightlight changes.
//
//	red:   only present in original
//	dim:   present in both
//	green: only present in new
func Diff(original, new Action) Action { _ = "STUB: not implemented"; return *new(Action) }

// TODO experimental - needs different signature
// TODO return patch format
func DiffPatch(original, new Action, c Context) []string { _ = "STUB: not implemented"; return nil }

package envsubst

// Eval replaces ${var} in the string based on the mapping function.
func Eval(s string, mapping func(string) string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// EvalEnv replaces ${var} in the string according to the values of the
// current environment variables. References to undefined variables are
// replaced by the empty string.
func EvalEnv(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

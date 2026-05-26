package common

type Messages struct {
	messages map[string]bool
}

func (m *Messages) init() {
	if m.messages == nil {
		m.messages = make(map[string]bool)
	}
}

func (m Messages) IsEmpty() bool {
	_ = "STUB: not implemented"
	// TODO replacement for Action.skipCache - does this need to consider suppressed messages or is this fine?
	return false
}

func (m *Messages) Add(s string) { _ = "STUB: not implemented"; return }

func (m Messages) Get() []string { _ = "STUB: not implemented"; return nil }

func (m *Messages) Suppress(expr ...string) error { _ = "STUB: not implemented"; return nil }

func (m *Messages) Merge(other Messages) { _ = "STUB: not implemented"; return }

func (m Messages) Integrate(values RawValues, prefix string) RawValues {
	_ = "STUB: not implemented"
	return *new(RawValues)
}

func (m Messages) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Messages) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

package common

type Meta struct {
	Messages Messages      `json:"messages"`
	Nospace  SuffixMatcher `json:"nospace"`
	Usage    string        `json:"usage"`
	Queries  Queries       `json:"queries,omitempty"`
}

func (m *Meta) Merge(other Meta) { _ = "STUB: not implemented"; return }

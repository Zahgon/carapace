package common

type SuffixMatcher struct {
	string
}

func (sm *SuffixMatcher) Add(suffixes ...rune) { _ = "STUB: not implemented"; return }

func (sm *SuffixMatcher) Merge(other SuffixMatcher) { _ = "STUB: not implemented"; return }

func (sm SuffixMatcher) Matches(s string) bool { _ = "STUB: not implemented"; return false }

func (sm SuffixMatcher) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sm *SuffixMatcher) UnmarshalJSON(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type ByRune []rune

func (r ByRune) Len() int           { _ = "STUB: not implemented"; return 0 }
func (r ByRune) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (r ByRune) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

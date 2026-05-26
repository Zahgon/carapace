package mock

type Mock struct {
	Dir     string
	Replies map[string]string
}

func (m Mock) CacheDir() string { _ = "STUB: not implemented"; return "" }

func (m Mock) WorkDir() string { _ = "STUB: not implemented"; return "" }

type t interface {
	Name() string
	Fatal(...any)
}

func NewMock(t t) *Mock { _ = "STUB: not implemented"; return nil }

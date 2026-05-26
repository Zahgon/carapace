package config

type configMap map[string]any

func (c configMap) Keys() []string { _ = "STUB: not implemented"; return nil }

type Field struct {
	Name        string
	Description string
	Style       string
	Tag         string
}

func (c configMap) Fields(name string) ([]Field, error) { _ = "STUB: not implemented"; return nil, nil }

var config = struct {
	Styles configMap
}{
	Styles: make(configMap),
}

func RegisterStyle(name string, i any) { _ = "STUB: not implemented"; return }

func Load() error { _ = "STUB: not implemented"; return nil }

func load(name string, c configMap) error { _ = "STUB: not implemented"; return nil }

func GetStyleConfigs() []string                   { _ = "STUB: not implemented"; return nil }
func GetStyleFields(name string) ([]Field, error) { _ = "STUB: not implemented"; return nil, nil }
func SetStyle(key, value string) error            { _ = "STUB: not implemented"; return nil }

func set(name, key, value string) error { _ = "STUB: not implemented"; return nil }

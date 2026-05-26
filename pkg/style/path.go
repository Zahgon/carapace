package style

type Context interface {
	Abs(s string) (string, error)
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
}

// ForPath returns the style for given path
//
//	/tmp/locally/reachable/file.txt
func ForPath(path string, sc Context) string { _ = "STUB: not implemented"; return "" }

// ForPath returns the style for given path by extension only
//
//	/tmp/non/existing/file.txt
func ForPathExt(path string, sc Context) string { _ = "STUB: not implemented"; return "" }

// ForExtension returns the style for given extension
//
//	json
func ForExtension(path string, sc Context) string { _ = "STUB: not implemented"; return "" }

func fromSGR(sgr string) string { _ = "STUB: not implemented"; return "" }

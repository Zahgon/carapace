//go:build linux

package ps

// Refresh reloads all the data associated with this process.
func (p *UnixProcess) Refresh() error { _ = "STUB: not implemented"; return nil }

// First, parse out the image name

// Move past the image name and start parsing the rest

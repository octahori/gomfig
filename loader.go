package gomfig

import "io"

// Loader defines the common behaviour for all configuration loaders
// (JSON, YAML, TOML, ENV, etc.). Implementations should be stateless so that
// they are safe for concurrent use.
//
// Callers typically use the convenience helpers in each format sub-package
// rather than interact with Loader directly.
type Loader interface {
	// Load decodes configuration data read from r into the value pointed to by
	// out.
	Load(r io.Reader, out interface{}) error

	// LoadFile opens the file located at path and delegates to Load.
	LoadFile(path string, out interface{}) error
}

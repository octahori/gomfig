package jsoncfg

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
)

// Loader implements gomfig.Loader for plain JSON documents.
//
// Zero-value is ready for use:
//
//    var cfg MyConfig
//    err := jsoncfg.Loader{}.LoadFile("config.json", &cfg)
//
type Loader struct{}

// Load decodes JSON data read from r into the value pointed to by out.
func (Loader) Load(r io.Reader, out interface{}) error {
    dec := json.NewDecoder(r)
    dec.DisallowUnknownFields()
    if err := dec.Decode(out); err != nil {
        return fmt.Errorf("gomfig/jsoncfg: decode json: %w", err)
    }
    return nil
}

// LoadFile opens path and delegates to Load.
func (l Loader) LoadFile(path string, out interface{}) error {
    f, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("gomfig/jsoncfg: open %s: %w", path, err)
    }
    defer f.Close()
    return l.Load(f, out)
}

// Convenience helpers -------------------------------------------------------

// LoadFile is a shorthand for Loader{}.LoadFile.
func LoadFile(path string, out interface{}) error { return (Loader{}).LoadFile(path, out) }

// Load is a shorthand for Loader{}.Load.
func Load(r io.Reader, out interface{}) error { return (Loader{}).Load(r, out) }

// MustLoadFile is like LoadFile but panics on error.
func MustLoadFile(path string, out interface{}) { if err := LoadFile(path, out); err != nil { panic(err) } }

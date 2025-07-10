package yamlcfg

import (
    "fmt"
    "io"
    "os"

    "gopkg.in/yaml.v3"
)

// Loader implements gomfig.Loader for YAML documents.
//
// YAML decoder is configured with KnownFields(true) to error on unknown fields,
// mirroring the strictness applied to JSON decoding.
//
// Usage:
//     var cfg MyConfig
//     err := yamlcfg.LoadFile("config.yaml", &cfg)
//
type Loader struct{}

// Load decodes YAML data from r into the value pointed to by out.
func (Loader) Load(r io.Reader, out interface{}) error {
    dec := yaml.NewDecoder(r)
    dec.KnownFields(true)
    if err := dec.Decode(out); err != nil {
        return fmt.Errorf("gomfig/yamlcfg: decode yaml: %w", err)
    }
    return nil
}

// LoadFile opens path and delegates to Load.
func (l Loader) LoadFile(path string, out interface{}) error {
    f, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("gomfig/yamlcfg: open %s: %w", path, err)
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

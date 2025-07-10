package tomlcfg

import (
    "fmt"
    "io"
    "os"

    "github.com/pelletier/go-toml/v2"
)

// Loader implements gomfig.Loader for TOML documents.
// It decodes strictly: unknown fields trigger an error via the UnmarshalOptions.
//
// Usage:
//     var cfg MyCfg
//     err := tomlcfg.LoadFile("config.toml", &cfg)
//
type Loader struct{}

// Load decodes TOML data from r into the value pointed to by out.
func (Loader) Load(r io.Reader, out interface{}) error {
    dec := toml.NewDecoder(r)
    dec.DisallowUnknownFields()
    if err := dec.Decode(out); err != nil {
        return fmt.Errorf("gomfig/tomlcfg: decode toml: %w", err)
    }
    return nil
}

// LoadFile opens path and delegates to Load.
func (l Loader) LoadFile(path string, out interface{}) error {
    f, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("gomfig/tomlcfg: open %s: %w", path, err)
    }
    defer f.Close()
    return l.Load(f, out)
}

// Convenience helpers -------------------------------------------------------

func LoadFile(path string, out interface{}) error { return (Loader{}).LoadFile(path, out) }
func Load(r io.Reader, out interface{}) error     { return (Loader{}).Load(r, out) }
func MustLoadFile(path string, out interface{})   { if err := LoadFile(path, out); err != nil { panic(err) } }

package gomfig

import (
    "io"

    "github.com/octahori/gomfig/yamlcfg"
)

// YAMLLoader aliases yamlcfg.Loader for convenience.
type YAMLLoader = yamlcfg.Loader

// LoadYAMLFile decodes YAML file into out.
func LoadYAMLFile(path string, out interface{}) error { return yamlcfg.LoadFile(path, out) }

// LoadYAML decodes YAML read from r into out.
func LoadYAML(r io.Reader, out interface{}) error { return yamlcfg.Load(r, out) }

// MustLoadYAMLFile is like LoadYAMLFile but panics on error.
func MustLoadYAMLFile(path string, out interface{}) { yamlcfg.MustLoadFile(path, out) }

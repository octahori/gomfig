package gomfig

import (
	"io"

	"github.com/octahori/gomfig/tomlcfg"
)

type TOMLLoader = tomlcfg.Loader

func LoadTOMLFile(path string, out interface{}) error { return tomlcfg.LoadFile(path, out) }
func LoadTOML(r io.Reader, out interface{}) error     { return tomlcfg.Load(r, out) }
func MustLoadTOMLFile(path string, out interface{})   { tomlcfg.MustLoadFile(path, out) }

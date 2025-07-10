package gomfig

// This file provides thin wrappers and type aliases so that users can keep
// importing the root package while the actual implementation lives in the
// `jsoncfg` sub-package. This avoids breaking changes as more loaders are
// added in future branches.

import (
    "io"

    "gomfig/jsoncfg"
)

// JSONLoader is kept for backwards-compatibility. It aliases jsoncfg.Loader.
type JSONLoader = jsoncfg.Loader

// LoadJSONFile calls jsoncfg.LoadFile.
func LoadJSONFile(path string, out interface{}) error { return jsoncfg.LoadFile(path, out) }

// LoadJSON calls jsoncfg.Load.
func LoadJSON(r io.Reader, out interface{}) error { return jsoncfg.Load(r, out) }

// MustLoadJSONFile calls jsoncfg.MustLoadFile.
func MustLoadJSONFile(path string, out interface{}) { jsoncfg.MustLoadFile(path, out) }

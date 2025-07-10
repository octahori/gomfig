package tomlcfg

import (
    "path/filepath"
    "testing"
    "os"
)

func writeTemp(t *testing.T, dir, name, content string) string {
    t.Helper()
    p := filepath.Join(dir, name)
    if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
        t.Fatalf("write temp: %v", err)
    }
    return p
}

func TestLoader_OK(t *testing.T) {
    tmp := t.TempDir()
    path := writeTemp(t, tmp, "cfg.toml", "[db]\nhost='localhost'\nport=1234\n")

    var cfg struct {
        DB struct {
            Host string `toml:"host"`
            Port int    `toml:"port"`
        } `toml:"db"`
    }

    if err := LoadFile(path, &cfg); err != nil {
        t.Fatalf("LoadFile error: %v", err)
    }
    if cfg.DB.Host != "localhost" || cfg.DB.Port != 1234 {
        t.Fatalf("unexpected result: %+v", cfg)
    }
}

func TestLoader_Unknown(t *testing.T) {
    tmp := t.TempDir()
    path := writeTemp(t, tmp, "bad.toml", "foo=1\n")
    var v struct{}
    if err := LoadFile(path, &v); err == nil {
        t.Fatal("expected error for unknown field")
    }
}

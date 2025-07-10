package yamlcfg

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTemp(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write temp file: %v", err)
	}
	return path
}

func TestLoader_LoadFile_Success(t *testing.T) {
	tmp := t.TempDir()
	path := writeTemp(t, tmp, "cfg.yaml", "db:\n  host: localhost\n  port: 1234\n")

	var cfg struct {
		DB struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"db"`
	}

	if err := LoadFile(path, &cfg); err != nil {
		t.Fatalf("LoadFile returned error: %v", err)
	}
	if cfg.DB.Host != "localhost" || cfg.DB.Port != 1234 {
		t.Fatalf("unexpected decoded data: %+v", cfg)
	}
}

func TestLoader_UnknownField(t *testing.T) {
	tmp := t.TempDir()
	path := writeTemp(t, tmp, "bad.yaml", "unknown: 42\n")

	var v struct{}
	if err := LoadFile(path, &v); err == nil {
		t.Fatal("expected error for unknown field, got nil")
	}
}

func TestLoader_MalformedYAML(t *testing.T) {
	tmp := t.TempDir()
	path := writeTemp(t, tmp, "broken.yaml", "foo: [bar\n")

	var v map[string]interface{}
	if err := LoadFile(path, &v); err == nil {
		t.Fatal("expected error for malformed yaml, got nil")
	}
}

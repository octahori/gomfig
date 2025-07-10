# gomfig  ![Version](https://img.shields.io/github/v/tag/octahori/gomfig?label=version)

> Current stable: **v0.2.0** (JSON, YAML, TOML)

Light-weight, extensible Go library to load and merge configuration from multiple sources.

`gomfig` is released under semantic versioning. Version **v0.2.0** supports **JSON, YAML, and TOML** configuration files. Upcoming minor versions will introduce ENV and remote back-ends. Upcoming branches will add YAML, TOML, ENV, and remote back-ends (Consul, etcd, …).

## Features (current branch)

* Decode JSON, YAML, and TOML files/streams into arbitrary Go structs or maps.
* Strict mode for all loaders: unknown fields are rejected early.
* Minimal API:

```go
// root package convenience helpers
type JSONLoader = gomfig.JSONLoader

var cfg struct {
    DB struct {
        Host string `json:"host"`
        Port int    `json:"port"`
    } `json:"db"`
}

if err := gomfig.LoadJSONFile("config.json", &cfg); err != nil {
    log.Fatalf("load config: %v", err)
}
```

## Installation

```sh
go get github.com/octahori/gomfig@v0.2.0
```

## Roadmap

| Stage | Description |
|-------|-------------|
| ✓ JSON | Basic loader implementation + unit tests |
| ✓ YAML | YAML loader (`yamlcfg`, strict mode) |
| ✓ TOML | TOML loader (`tomlcfg`) |
| ☐ ENV  | Map environment variables into struct with tags |
| ☐ Remote | Consul / etcd loaders |

## Contributing

1. Fork and create a feature branch.
2. Ensure `go test ./...` and `go vet ./...` pass.
3. Open a pull request.

---

Licensed under the MIT License.
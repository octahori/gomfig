# gomfig

Light-weight, extensible Go library to load and merge configuration from multiple sources.

This **json-loader** branch focuses on JSON support only. Upcoming branches will add YAML, TOML, ENV, and remote back-ends (Consul, etcd, …).

## Features (current branch)

* Decode JSON files/streams into arbitrary Go structs or maps.
* Strict mode: unknown fields are rejected early (`json.Decoder.DisallowUnknownFields`).
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
go get github.com/your-user/gomfig@latest
```

## Roadmap

| Stage | Description |
|-------|-------------|
| ✓ JSON | Basic loader implementation + unit tests |
| ☐ YAML | `yamlcfg` sub-package using `gopkg.in/yaml.v3` |
| ☐ TOML | `tomlcfg` sub-package using `github.com/pelletier/go-toml/v2` |
| ☐ ENV  | Map environment variables into struct with tags |
| ☐ Remote | Consul / etcd loaders |

## Contributing

1. Fork and create a feature branch.
2. Ensure `go test ./...` and `go vet ./...` pass.
3. Open a pull request.

---

Licensed under the MIT License.
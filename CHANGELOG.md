# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.2.2] - 2025-07-10
### Fixed
- Documentation updates: README stable version badge and install instructions.

## [v0.2.0] - 2025-07-10
### Added
- TOML loader (`tomlcfg`) with strict decoding.
- Root helpers: `LoadTOMLFile`, `MustLoadTOMLFile`, etc.
- README roadmap update and example.

### Fixed
- `gofmt` compliance for new TOML files.

## [v0.1.0] - 2025-07-10
### Added
- Initial release with JSON configuration loader (`jsoncfg`).
- Strict decoding (`DisallowUnknownFields`).
- Public API helpers: `LoadJSONFile`, `MustLoadJSONFile`, etc.
- GitHub Actions CI (build, vet, test, gofmt).
- README with usage, roadmap, contribution guide.

### Removed
- Legacy `json_loader.go` file (replaced by `jsoncfg`).

[Unreleased]: https://github.com/octahori/gomfig/compare/v0.2.2...HEAD
[v0.2.2]: https://github.com/octahori/gomfig/releases/tag/v0.2.2
[v0.2.0]: https://github.com/octahori/gomfig/releases/tag/v0.2.0
[v0.1.0]: https://github.com/octahori/gomfig/releases/tag/v0.1.0

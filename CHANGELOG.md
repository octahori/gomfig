# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.1.0] - 2025-07-10
### Added
- Initial release with JSON configuration loader (`jsoncfg`).
- Strict decoding (`DisallowUnknownFields`).
- Public API helpers: `LoadJSONFile`, `MustLoadJSONFile`, etc.
- GitHub Actions CI (build, vet, test, gofmt).
- README with usage, roadmap, contribution guide.

### Removed
- Legacy `json_loader.go` file (replaced by `jsoncfg`).

[Unreleased]: https://github.com/octahori/gomfig/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/octahori/gomfig/releases/tag/v0.1.0

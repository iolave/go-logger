# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v1.0.0] - 2025-12-13

### Added
- Automatic inclusion of trace information from context via [iolave/go-trace](https://github.com/iolave/go-trace).
- Circular references handling in log data.
- Trace log level.
- Trace and TraceWithData methods.

### Changed
- Schema version v1.0.0 structure.
- Log level definition.
- Logger is now an interface.
- Logger methods signature.

## [v0.1.1] - 2024-07-28

### Fixed
- Time property. It showed as string despite being changed to int internally.

## [v0.1.0] - 2024-07-28

### Added

- LOG_LEVEL environment variable support.
- LOG_LEVEL constants.
- Debug, Info, Warn, Error and Fatal log methods.
- String helper method that converts strings to their snake case version.

[Unreleased]: https://github.com/iolave/go-logger/compare/v1.0.0...HEAD
[v1.0.0]: https://github.com/iolave/go-logger/compare/v0.1.1...v1.0.0
[v0.1.1]: https://github.com/iolave/go-logger/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/iolave/go-logger/releases/tag/v0.1.0

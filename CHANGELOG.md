# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Windows ARM64 shared library (`a2l_grpc_windows_arm64.dll`), built and tested natively in CI.

## [0.2.0] - 2026-07-31

### Added

- Test suite covering the A2L grammar tokens.

### Changed

- Aligned the A2L grammar with the ASAP2 1.6.1 specification: added `CALIBRATION_HANDLE_TEXT`,
  made `MATRIX_DIM` require all three dimensions, and used floating point values for
  `EXTENDED_LIMITS` and `STEP_SIZE`.

### Removed

- Keywords which are not part of the ASAP2 1.6.1 specification: `INSTANCE`, `TYPEDEF_CHARACTERISTIC`,
  `TYPEDEF_MEASUREMENT`, `TYPEDEF_STRUCTURE`, `STRUCTURE_COMPONENT`, `LINK_TYPE` and
  `ALIGNMENT_FLOAT16_IEEE`.

## [0.1.22] - 2026-07-30

### Added

- macOS Intel shared library.

### Fixed

- A panic raised in the Go library no longer terminates the calling process ([#12](https://github.com/Sauci/a2l-grpc/issues/12)).
- Linux ARM64, Linux i386 and 32-bit Windows builds.

## [0.1.21] - 2025-02-28

### Fixed

- Architecture name used in the released artifacts.

## [0.1.20] - 2025-02-28

### Added

- Linux i386 shared library.

## [0.1.19] - 2025-01-28

### Added

- Maximum gRPC message size as an argument of the shared library entry point.

## [0.1.18] - 2025-01-27

### Added

- Streaming gRPC API, allowing A2L files larger than the maximum message size to be exchanged.
- `COEFFS_LINEAR` support for `COMPU_METHOD`.
- Base functions exported by the shared library.

### Changed

- Updated Go to 1.22 and protoc to 29.3.

## [0.1.17] - 2024-04-12

### Added

- Linux ARM and ARM64 shared libraries.

## [0.1.16] - 2024-02-16

### Added

- 32-bit Windows shared library.

### Fixed

- Restored the macOS build.

## [0.1.15] - 2023-12-13

### Added

- macOS ARM64 shared library.

## [0.1.14] - 2023-10-01

### Added

- `DISCRETE`, `SYMBOL_LINK`, `STEP_SIZE`, `PHYS_UNIT` and `LAYOUT` support.

## [0.1.13] - 2023-09-17

### Changed

- Sorted nodes naturally, so that `"2"` comes before `"10"`.

## [0.1.12] - 2023-08-30

### Added

- Optional sorting of the `MODULE` node children.

## [0.1.11] - 2023-07-04

### Changed

- Unified all data transfer formats of the gRPC API.

## [0.1.10] - 2023-07-04

### Fixed

- Indentation of the generated A2L file.

## [0.1.9] - 2023-06-30

### Fixed

- Encoding issue observed with Python versions older than 3.10.

## [0.1.8] - 2023-06-30

### Changed

- Improved error reporting of the parsing methods.

## [0.1.7] - 2023-06-30

### Added

- Conversion from a tree to an A2L file, with a configurable indentation.

### Fixed

- Integer and long values are dumped according to their original representation.

## [0.1.6] - 2023-06-28

### Added

- Partial JSON inputs and outputs support in the gRPC interface.

## [0.1.5] - 2023-06-18

### Fixed

- Handling of a `0` indentation value.

## [0.1.4] - 2023-06-16

### Added

- Conversion between a tree and its JSON representation.

### Changed

- Cleaned up the gRPC API.

## [0.1.3] - 2023-06-13

### Fixed

- Escaped double quotes in A2ML keyword values.

## [0.1.2] - 2023-06-13

### Fixed

- Dump of empty messages ([protobuf#2954](https://github.com/protocolbuffers/protobuf/issues/2954)).

## [0.1.1] - 2023-06-13

### Added

- `Present` field in empty messages, so that their presence can be checked on the receiver side.

### Fixed

- `ANNOTATION` grammar, `AXIS_PTS_X/Y/Z` properties, `GROUP` parsing and `DATA_SIZE` type format.

## [0.1.0] - 2023-06-08

### Added

- Initial release: ANTLR based A2L parser exposed through a gRPC API, shipped as a shared library
  for Linux and Windows, together with the protobuf definitions.

[Unreleased]: https://github.com/Sauci/a2l-grpc/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/Sauci/a2l-grpc/compare/v0.1.22...v0.2.0
[0.1.22]: https://github.com/Sauci/a2l-grpc/compare/v0.1.21...v0.1.22
[0.1.21]: https://github.com/Sauci/a2l-grpc/compare/v0.1.20...v0.1.21
[0.1.20]: https://github.com/Sauci/a2l-grpc/compare/v0.1.19...v0.1.20
[0.1.19]: https://github.com/Sauci/a2l-grpc/compare/v0.1.18...v0.1.19
[0.1.18]: https://github.com/Sauci/a2l-grpc/compare/v0.1.17...v0.1.18
[0.1.17]: https://github.com/Sauci/a2l-grpc/compare/v0.1.16...v0.1.17
[0.1.16]: https://github.com/Sauci/a2l-grpc/compare/v0.1.15...v0.1.16
[0.1.15]: https://github.com/Sauci/a2l-grpc/compare/v0.1.14...v0.1.15
[0.1.14]: https://github.com/Sauci/a2l-grpc/compare/v0.1.13...v0.1.14
[0.1.13]: https://github.com/Sauci/a2l-grpc/compare/v0.1.12...v0.1.13
[0.1.12]: https://github.com/Sauci/a2l-grpc/compare/v0.1.11...v0.1.12
[0.1.11]: https://github.com/Sauci/a2l-grpc/compare/v0.1.10...v0.1.11
[0.1.10]: https://github.com/Sauci/a2l-grpc/compare/v0.1.9...v0.1.10
[0.1.9]: https://github.com/Sauci/a2l-grpc/compare/v0.1.8...v0.1.9
[0.1.8]: https://github.com/Sauci/a2l-grpc/compare/v0.1.7...v0.1.8
[0.1.7]: https://github.com/Sauci/a2l-grpc/compare/v0.1.6...v0.1.7
[0.1.6]: https://github.com/Sauci/a2l-grpc/compare/v0.1.5...v0.1.6
[0.1.5]: https://github.com/Sauci/a2l-grpc/compare/v0.1.4...v0.1.5
[0.1.4]: https://github.com/Sauci/a2l-grpc/compare/v0.1.3...v0.1.4
[0.1.3]: https://github.com/Sauci/a2l-grpc/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/Sauci/a2l-grpc/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/Sauci/a2l-grpc/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/Sauci/a2l-grpc/releases/tag/v0.1.0

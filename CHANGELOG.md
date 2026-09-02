# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Windows ARM64 shared library (`a2l_grpc_windows_arm64.dll`), built and tested natively in CI.

### Added

- Report of the keywords which ASAM MCD-2 MC 1.6.1 withdrew (chapter 1.4.4): `S_REC_LAYOUT`,
  `NO_RESCALE_Y/_Z/_4/_5` and `AXIS_RESCALE_Y/_Z/_4/_5`. They stay accepted, the grammar also
  covers ASAP2 1.51, but a file declaring 1.61 or newer is now warned about them.
- Report of a keyword used more than once where the specification allows a single occurrence
  (chapter 3.5.1). Such a repeat parsed silently before and only the last occurrence was kept.
- Report of a missing `ASAP2_VERSION` when `EnforceVersionCheck` is set. The keyword is mandatory
  since 1.6.1 (chapters 1.4.4 and 3.5.16), and without it no version gate can run at all.

### Fixed

- A2L files starting with a UTF-8 byte-order mark are now parsed instead of being rejected on
  their first character. The mark declares the encoding of the file (ASAM MCD-2 MC 1.6.1,
  chapter 1.5.2) and is not part of its content.
- `ALPHA` is usable as an identifier again. Making it a `VAR_NAMING` keyword reserved the word
  everywhere, which rejected files the specification itself shows: chapter 3.5.29 uses `ALPHA` as
  an identifier in its own example. `VAR_NAMING ALPHA` is still accepted and still warned about.
- `BIT_MASK` and `ERROR_MASK` accept the whole unsigned 64 bit range, as required since 1.6.1
  (chapter 1.4.4). A mask above the signed range previously failed to parse.
- A string containing an escape sequence the specification does not define is reported on its own
  instead of desynchronizing the lexer. A single unescaped Windows path used to produce a cascade
  of unrelated errors, the last of which swallowed the remainder of the file.
- `HEADER` may appear anywhere among the `MODULE` blocks of a `PROJECT`. Neither version
  prescribes an order for the optional elements; at most one `HEADER` is still enforced.
- An identifier is decomposed into its partial identifiers as chapter 3.2 describes, so a
  malformed one such as `a.`, `a..b` or `a.0b` is rejected instead of being taken as a single
  identifier. A signed or dotted array index, e.g. `a[-1]` or `a[x.y]`, is rejected as well.

### Changed

- The `Mask` of `BIT_MASK` and `ERROR_MASK` is a `ULongType` instead of a `LongType`, so that it
  can carry the unsigned 64 bit range. **This changes the protobuf message shape.**
- The labels of the grammar now match the cardinality of the specification, and the parser uses
  them to detect a repeated single occurrence keyword.

## [0.2.0] - 2026-07-31

### Added

- Test suite covering the A2L grammar tokens.

### Changed

- Aligned the A2L grammar with the ASAP2 1.6.1 specification: added `CALIBRATION_HANDLE_TEXT`,
  allowed `IF_DATA` in `GROUP`, made `MATRIX_DIM` require all three dimensions, and anchored the
  end of the file so that content behind `/end PROJECT` is no longer ignored.
- Every parameter declared as `float` by the specification now accepts a floating point value:
  `EXTENDED_LIMITS`, `STEP_SIZE`, `MAX_GRAD`, `COEFFS`, `COEFFS_LINEAR`, `UNIT_CONVERSION`,
  `DEFAULT_VALUE_NUMERIC`, the `FIX_AXIS_PAR`, `FIX_AXIS_PAR_DIST` and `FIX_AXIS_PAR_LIST`
  parameters, the value lists of `COMPU_TAB`, `COMPU_VTAB` and `COMPU_VTAB_RANGE`, the limits and
  `MaxDiff` of `AXIS_PTS`, `CHARACTERISTIC` and `AXIS_DESCR`, and the accuracy and limits of
  `MEASUREMENT`. A hexadecimal value is no longer accepted for such a parameter.
- **Breaking:** only the escape sequences listed by the specification (`\"`, `\'`, `\\`, `\n`,
  `\r` and `\t`, plus a doubled `""`) are accepted inside a string. A backslash followed by any
  other character, e.g. an unescaped Windows path such as `"C:\data\file"`, is now rejected.
- **Breaking:** `A2ML` must be declared directly after the long identifier of `MODULE`, as
  required by chapter 3.5.89, and may appear only once.
- **Breaking:** `MATRIX_DIM` now reports all three dimensions in the tree, so consumers reading
  only `XDim` see a changed message shape.

### Removed

- Keywords which are not part of the ASAP2 1.6.1 specification: `INSTANCE`, `TYPEDEF_CHARACTERISTIC`,
  `TYPEDEF_MEASUREMENT`, `TYPEDEF_STRUCTURE`, `STRUCTURE_COMPONENT`, `LINK_TYPE` and
  `ALIGNMENT_FLOAT16_IEEE`.
- The `FLOAT16_IEEE` data type, and the `int64` and `uint64` predefined type names of the A2ML
  metalanguage. All three were introduced by ASAM MCD-2 MC 1.7 and appear in neither ASAP2 1.51
  nor 1.6.1.

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

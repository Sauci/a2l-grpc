# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2026-09-03

### Added

- Report of an `/include` statement, naming the file it refers to. The mechanism is a text
  replacement which needs the file system the file was read from (chapter 4), so a parser working
  on the content of a single file cannot resolve it; substituting the included files is left to the
  caller. Such a file used to be rejected with a syntax error naming neither the statement nor the
  file, and chapter 4 places no restriction on where the statement may appear.
- Report of a second `CALIBRATION_HANDLE` in a `CALIBRATION_METHOD` when the file declares 1.61 or
  newer, which reduced it to a single occurrence (chapter 3.5.28). ASAP2 1.51 (chapter 6.3.26)
  allows several, so the grammar keeps accepting them.
- `go vet` runs in CI.
- Report of the keywords which ASAM MCD-2 MC 1.6.1 withdrew (chapter 1.4.4): `S_REC_LAYOUT`,
  `NO_RESCALE_Y/_Z/_4/_5` and `AXIS_RESCALE_Y/_Z/_4/_5`. They stay accepted, the grammar also
  covers ASAP2 1.51, but a file declaring 1.61 or newer is now warned about them.
- Report of a keyword used more than once where the specification allows a single occurrence
  (chapter 3.5.1). Such a repeat parsed silently before and only the last occurrence was kept.
- Report of a missing `ASAP2_VERSION` when `EnforceVersionCheck` is set. The keyword is mandatory
  since 1.6.1 (chapters 1.4.4 and 3.5.16), and without it no version gate can run at all.

### Changed

- **Breaking:** the gRPC server is bound to the loopback interface instead of every interface. It
  is an unauthenticated endpoint serving the process which loaded the library, and the documented
  usage connects to `localhost`; it used to be reachable from the whole network.
- **Breaking:** a taggedstruct member must carry its tag, which chapter 5.2 makes mandatory in both
  forms of `taggedstruct_definition`. A tagless member parsed before and then made the serializer
  dereference a nil tag.
- **Breaking:** the array specifier of A2ML takes a constant, as chapter 5.2 declares it. It used
  to share the rule of the A2L grammar, which also accepts an alphabetic string, because chapter
  3.2 allows one as the index of a partial identifier; a symbolic size reached an unimplemented
  conversion and was reported as an internal error.
- The `Mask` of `BIT_MASK` and `ERROR_MASK` is a `ULongType` instead of a `LongType`, so that it
  can carry the unsigned 64 bit range. **This changes the protobuf message shape.**
- The labels of the grammar now match the cardinality of the specification, and the parser uses
  them to detect a repeated single occurrence keyword.
- `StringType` documents that it carries the string as written in the file: the escape sequences of
  chapter 3.2 are left in place, so that the value is reproduced verbatim when the tree is
  serialized back to A2L.
- The ANTLR generator is pinned to 4.13.2 and verified against its checksum, instead of being built
  from a clone of the default branch of the upstream repository on every run. `protoc` is pinned
  and verified the same way.
- CI no longer runs `go get`, which resolved the ANTLR runtime to its latest version and rewrote
  `go.mod`, so the dependency versions the module declares are the ones tested and released. Every
  job which runs Go now installs the toolchain `go.mod` declares, instead of taking the one which
  happens to sit on the runner.
- **Breaking:** the module requires Go 1.25, which is the version `google.golang.org/grpc` needs.
  The toolchain it is built with is 1.27.1. Besides the newer dependencies this is worth having on
  its own: measured on a 1.3 MB file, the compiler of 1.27 parses about a quarter faster than the
  one of 1.22 (2.88 s to 2.19 s) and removes about a fifth of the allocations (39.2 to 31.0
  million), without a change to this repository.
- Updated the toolchain and the dependencies: Go 1.22 to 1.25/1.27.1, ANTLR 4.13.1 to 4.13.2 with
  its runtime `github.com/antlr4-go/antlr/v4` v4.13.0 to v4.13.1, which the generated parser has to
  match, `google.golang.org/grpc` v1.69.4 to v1.83.2, `google.golang.org/protobuf` v1.35.1 to
  v1.36.12, `github.com/stretchr/testify` v1.8.4 to v1.12.1, `protoc` 29.3 to 36.1,
  `protoc-gen-go` v1.33.0 to v1.36.12 and `protoc-gen-go-grpc` v1.3.0 to v1.6.2. The protobuf
  messages are unchanged, and so are the `.proto` files a client generates its own sources from.

### Fixed

- No response of the gRPC server exceeds the maximum message size any more. The warnings all rode
  on the first response, on top of a chunk of tree sized to fill it, and an error was sent in one
  piece whatever its length; a large file with many warnings, or a badly broken one, made gRPC
  reject the response with `RESOURCE_EXHAUSTED`, a transport error which named neither the file nor
  the cause, and which raising the limit on both sides merely postponed. The warnings now come in
  responses of their own, before the chunks of the tree, spread over as many as they need, and an
  error longer than the limit is shortened to its first lines and says how many were left out.
  **The warnings are no longer on the first response** but on the leading ones; a client which
  collects them from every response, as the one of the README does, is unaffected.
- A partial identifier may begin with a digit, e.g. `SFB_R_FFO_DE.Properties.1.Qly`. Chapter 3.2
  lists two limitations for an identifier: "the first character must be a letter or an underscore,
  brackets must occur in pairs at the end of a partial string". The first one is about the first
  character of the identifier, not of every partial string; where the chapter means a partial
  string it says so, as the second limitation does. Such identifiers occur in the wild
  ([pya2l#17](https://github.com/Sauci/pya2l/issues/17)) and were rejected. An empty partial
  string, as in `a.` or `a..b`, is still rejected, and so is a symbolic array index containing a
  point, which chapter 3.2 describes as a single enumerator of the C program.
- The dimension of an A2ML member of a named type, e.g. `struct my_type[4];`, is reported. The
  brackets were consumed as the array index of the identifier, which left the member without its
  `ArraySpecifier`.
- An identifier of the metalanguage may be spelled like an A2L keyword, as chapter 5.2 requires:
  "Within the AML own name spaces are used. In this case it is allowed to reuse ASAM MCD-2 MC
  keyword names." The A2L, A2ML and IF_DATA grammars share one lexer, in which every keyword is a
  token of its own, so such an identifier never reached the rule for an identifier. The reference
  AML of the specification, which declares `taggedunion IF_DATA`, was rejected; so was an IF_DATA
  blob containing an element named after an A2L keyword. Chapter 3.2 still forbids them outside of
  A2ML and IF_DATA.
- The `tag "(" member ")*;"` form of a taggedstruct member (chapter 5.2) is accepted. Its tag was
  matched against a token which the shared lexer can never produce, so the form was unparsable.
- An `int` or `uint` parameter whose value does not fit is reported, with its position, instead of
  being truncated. `ALIGNMENT_BYTE 0xFFFFFFFF` used to be kept as `-1` and dumped as
  `0x-0000001`, which no longer parses.
- The messages of the parser are no longer used as a format string, so a per cent sign in the
  parsed content does not reach the caller as `%!'(NOVERB)` any more. The indentation of a `block`
  of the metalanguage was formatted the same way when a tree was serialized back to A2L.
- `Create` returns 1 when the server could not be started. It used to return 0 whenever no server
  was running yet, the port could not be bound included, so a caller checking the result went on to
  connect to a port nothing was listening on.
- `Create` rejects a maximum message size which leaves no room for a chunk of tree. A size below
  the protocol margin used to make the response loop spin forever or slice backwards.
- `GetA2LFromTree` reports a tree it cannot serialize through the `error` field of the response,
  like the other methods, instead of losing the stream to an internal error. A tree which lacks an
  element the parser always fills, but which a client building a tree by hand may leave out, made
  the serializer dereference a nil node.
- Nodes which share a sort key keep the order of the file, instead of being shuffled from one run
  to the next.
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

## [0.2.1] - 2026-07-31

### Added

- Windows ARM64 shared library (`a2l_grpc_windows_arm64.dll`), built and tested natively in CI.

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

[Unreleased]: https://github.com/Sauci/a2l-grpc/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/Sauci/a2l-grpc/compare/v0.2.1...v0.3.0
[0.2.1]: https://github.com/Sauci/a2l-grpc/compare/v0.2.0...v0.2.1
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

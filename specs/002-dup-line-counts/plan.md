# Implementation Plan: Duplicate Line Counts

**Branch**: `main` | **Date**: 2026-11-19 | **Spec**: [spec.md](spec.md)

**Input**: Feature specification for `002-dup-line-counts`.

## Summary

Make `ch1/dup` a runnable Chapter 1 example that reads lines from standard
input, counts identical lines, and writes only duplicates (count greater than
one) to standard output as `<count><tab><line>`. Extract the counting and
reporting logic from the entry point so it can be unit-tested with injected
readers and writers, matching the constitution's package test gate.

## Technical Context

**Language/Version**: Go 1.26 (module minimum).

**Primary Dependencies**: Go standard library for the command; Testify for unit tests only.

**Storage**: None; line counts are held in memory for one run and discarded.

**Testing**: Standard-library `testing` with Testify assertions in
`ch1/dup/dup_test.go`. Cover the duplicate-detection cases in
[VR-001](spec.md#verification-requirements). Assert the set of reported
line/count pairs (order unspecified). Run `go test ./ch1/dup`, `gofmt -l` on
changed Go files, and the CLI checks in [quickstart.md](quickstart.md).

**Target Platform**: Local command line with the Go toolchain.

**Project Type**: Standalone command-line learning example.

**Performance Goals**: Process typical textbook-scale inputs (dozens of lines)
comfortably within the one-minute learner check in [SC-002](spec.md#measurable-outcomes);
no throughput target.

**Constraints**: Read standard input only (no file arguments). Report only lines
with count greater than one. Preserve exact line text, including empty lines and
whitespace. No new runtime dependencies or changes to other examples.

**Scale/Scope**: One executable example under `ch1/dup`.

## Constitution Check

| Principle | Planning gate |
|-----------|---------------|
| Learning Scope | Pass: confined to the Chapter 1 duplicate-finding example. |
| Exercise Fidelity | Pass: stdin-only duplicate reporting matches the chapter's first `dup` form; file arguments are explicitly out of scope and not presented as the fuller book variants. |
| Go Formatting | Pass for design: all added or modified Go files must be formatted. |
| Unit Tests | Pass for design: counting/reporting behavior moves out of a `main`-only file into a tested source file; the `main`-only entry point remains exempt; all tests in `ch1/dup` must pass. |

Phase 1 re-check: the data model and CLI contract stay stdin-only with unordered
duplicate reports. The validation guide checks formatting, the affected package's
unit tests, and CLI behavior against the contract.

## Project Structure

### Documentation (this feature)

```text
specs/002-dup-line-counts/
├── spec.md
├── plan.md
├── research.md
├── data-model.md
├── quickstart.md
└── contracts/
    └── cli.md
```

### Source Code (repository root)

```text
ch1/
└── dup/
    ├── main.go      # Program entry point (stdin → stdout wiring)
    ├── dup.go       # Line counting and duplicate reporting
    └── dup_test.go  # Behavior unit tests
```

**Structure Decision**: Keep `main.go` as the entry point that wires
`os.Stdin`/`os.Stdout` into package functions defined in `dup.go`. Give the
shared logic an `io.Reader` and `io.Writer` so tests can supply buffers without
redirecting process-wide standard streams. Reuse or replace the existing
`ch1/dup/main.go` scaffolding so the package remains runnable with
`go run ./ch1/dup`.

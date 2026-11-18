# Implementation Plan: Echo Arguments with Loops

**Branch**: `main` | **Date**: 2026-11-18 | **Spec**: [spec.md](spec.md)

**Input**: Feature specification for `001-hello-world-loops`.

## Summary

Update the existing `ch1/echo` command so learners can pass arguments
at launch. Keep the four existing loop forms as independent demonstrations. Each
receives the same argument list, builds the same space-separated text, and emits
one newline-terminated copy. Keep the non-loop demonstration, separators, and
`END` output removed. Add unit tests that check the exact output of each loop.

## Technical Context

**Language/Version**: Go 1.26 (module minimum).

**Primary Dependencies**: Go standard library for the command; Testify for unit tests only.

**Storage**: None; arguments and results are transient.

**Testing**: Standard-library `testing` with Testify assertions in
`ch1/echo/loops_test.go`. Check each loop against the four input cases in
[VR-001](spec.md#verification-requirements).
Run `go test ./ch1/echo`, `gofmt -l` on changed Go files, and the CLI checks in
[quickstart.md](quickstart.md).

**Target Platform**: Local command line with the Go toolchain.

**Project Type**: Standalone command-line learning example.

**Performance Goals**: All four demonstrations terminate without interactive input;
no throughput or latency target.

**Constraints**: Preserve argument order and content, place one space between
arguments, and print exactly four copies. No new runtime dependencies or changes
to other examples.

**Scale/Scope**: One executable example under `ch1/echo`.

## Constitution Check

| Principle | Planning gate |
|-----------|---------------|
| Learning Scope | Pass: confined to a Chapter 1 learning example demonstrating four Go loop forms. |
| Exercise Fidelity | Pass: the spec identifies the example as a learning variant, not the book's original solution. |
| Go Formatting | Pass for design: all added or modified Go files must be formatted. |
| Unit Tests | Pass for design: `loops.go` has planned behavior tests, the `main`-only entry point is exempt, and all tests in `ch1/echo` must pass. |

Phase 1 re-check: the argument model and CLI contract are unchanged. The validation
guide now checks formatting, the affected package's unit tests, and CLI behavior.

## Project Structure

### Documentation (this feature)

```text
specs/001-hello-world-loops/
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
└── echo/
    ├── main.go        # Program entry point
    ├── loops.go       # Four loop demonstrations
    └── loops_test.go  # Exact-output unit tests
```

**Structure Decision**: Keep `main.go` as the entry point. It reads the arguments
once, excluding the program name, and calls the four demonstrations defined in
`loops.go` in the same package. Each loop builds its result independently, using
argument position to decide when to add a separator. Give each demonstration an
`io.Writer` so the entry point can pass standard output and tests can pass a
buffer. Return write errors to the entry point. Run the entire package with
`go run ./ch1/echo`.

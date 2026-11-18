<!--
Sync Impact Report
Version change: none (1.0.0; completes the initial constitution)
Principles added: Unit Tests
Principle modified: Formatting and Tests -> Go Formatting
Section updated: Review Evidence
Sections added or removed: none
Follow-up: Align the active echo feature's plan, tasks, and validation guide
with the package-scoped test gate before resuming implementation.
-->

# gopl Constitution

## Core Principles

### I. Learning Scope

- **Rule:** Each change MUST advance an identifiable chapter, exercise, or Go concept, or fix the correctness or runnability of an existing example.
- **Reject if:** A plan adds an unrelated web admin interface to a command-line exercise; it serves no learning objective.

### II. Exercise Fidelity

- **Rule:** Exercise solutions MUST satisfy the exercise's explicit constraints. Any departure MUST be labeled as a variant and MUST NOT be presented as the original solution.
- **Reject if:** A PR uses a library function where the exercise explicitly requires implementing the algorithm by hand and presents it as the original solution.

### III. Go Formatting

- **Rule:** All added or modified Go files MUST be formatted with `gofmt`.
- **Reject if:** `gofmt -l` lists a Go file added or modified by the PR.

### IV. Unit Tests

- **Rule:** Every Go source file other than a `*_test.go` file MUST have unit tests covering its behavior, except files that define only `main` or only constants. Every unit test in each affected Go package MUST pass; unrelated packages are not part of this gate.
- **Reject if:** A PR adds `ch1/echo/loops.go` without unit tests covering its loop behavior, or any unit test in `ch1/echo` fails.

## Project Scope

This repository is for studying [The Go Programming Language](https://golang-china.github.io/gopl-zh/).
It is not a production framework.

## Review Evidence

Review plans for exercise constraints and required unit tests. For PRs, check formatting
and test coverage for changed Go source files, then run `go test` on each affected package
(for example, `go test ./ch1/echo`). A package reporting `[no test files]` does not satisfy
the unit-test requirement if it contains a non-exempt file that needs coverage.

## Governance

Reject any plan or PR that violates a principle. To change a rule, approve a separate
constitution amendment first; no case-by-case waivers. Use a major version bump for
incompatible changes, minor for new or expanded rules, and patch for wording-only changes.

**Version**: 1.0.0 | **Ratified**: 2026-11-18 | **Last Amended**: 2026-11-18

<!--
Sync Impact Report
Version change: unratified scaffold -> 1.0.0 (initial adoption)
Principles added: Learning Scope; Exercise Fidelity; Formatting and Tests
Principles modified or removed: none
Sections added: Project Scope; Review Evidence (replacing generic template slots)
Sections removed: none; two unused principle slots omitted
Follow-up TODOs: none
-->

# gopl Constitution

## Core Principles

### I. Learning Scope

- **Rule:** Each change MUST advance an identifiable chapter, exercise, or Go concept, or fix the correctness or runnability of an existing example.
- **Reject if:** A plan adds an unrelated web admin interface to a command-line exercise; it serves no learning objective.

### II. Exercise Fidelity

- **Rule:** Exercise solutions MUST satisfy the exercise's explicit constraints. Any departure MUST be labeled as a variant and MUST NOT be presented as the original solution.
- **Reject if:** A PR uses a library function where the exercise explicitly requires implementing the algorithm by hand and presents it as the original solution.

### III. Formatting and Tests

- **Rule:** All added or modified Go files MUST be formatted with `gofmt`, and `go test ./...` MUST pass from the repository root.
- **Reject if:** A PR adds an exercise that causes any package to fail to compile under `go test ./...`.

## Project Scope

This repository is for studying [The Go Programming Language](https://golang-china.github.io/gopl-zh/).
It is not a production framework.

## Review Evidence

Review plans against the proposed changes and exercise constraints. For PRs, review the diff,
check that `gofmt -l` lists none of the added or modified Go files, and check `go test ./...`.

## Governance

Reject any plan or PR that violates a principle. To change a rule, approve a separate
constitution amendment first; no case-by-case waivers. Use a major version bump for
incompatible changes, minor for new or expanded rules, and patch for wording-only changes.

**Version**: 1.0.0 | **Ratified**: 2026-11-18 | **Last Amended**: 2026-11-18

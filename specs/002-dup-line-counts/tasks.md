# Tasks: Duplicate Line Counts

**Input**: [Implementation plan](plan.md), [specification](spec.md), [research](research.md),
[data model](data-model.md), [CLI contract](contracts/cli.md), and [quickstart](quickstart.md).

**Tests**: Add behavior unit tests for `ch1/dup/dup.go` per
[VR-001](spec.md#verification-requirements). `main.go` remains an entry-point-only
file and does not need dedicated unit tests.

## Phase 1: Setup

No tasks. Reuse the Go module declared in `go.mod` and the existing `ch1/dup` package.

## Phase 2: Foundational

No tasks. This standalone command needs no shared infrastructure or dependencies.

## Phase 3: User Story 1 - Find duplicate lines from standard input (Priority: P1) 🎯 MVP

**Goal**: Count identical input lines and emit only duplicate reports with their counts.

**Independent test**: Run `go test ./ch1/dup` and confirm cases for duplicates,
unique-only lines, empty input, empty duplicated lines, and whitespace-sensitive
identity. Assert the set of reported line/count pairs (order unspecified).

### Tests for User Story 1

> Write these tests first and confirm they fail against any incomplete
> counting/reporting API before finishing the implementation task.

- [X] T001 [P] [US1] Add table-driven tests in `ch1/dup/dup_test.go` for the
  counting/reporting API using an `io.Reader` / `io.Writer` and a fresh buffer
  per case. Cover duplicates, unique-only lines, empty input, empty duplicated
  lines, and whitespace-sensitive identity per
  [VR-001](spec.md#verification-requirements). Assert the set of
  [`Duplicate report.count`](data-model.md) /
  [`Duplicate report.line text`](data-model.md) pairs rather than byte order when
  multiple reports exist. Confirm the new tests cannot pass before
  `ch1/dup/dup.go` implements that API.

### Implementation for User Story 1

- [X] T002 [US1] Implement line counting and duplicate reporting in
  `ch1/dup/dup.go`. Read newline-delimited [`Input line.content`](data-model.md)
  values from an `io.Reader`, maintain [`Occurrence count.value`](data-model.md)
  per distinct line, and write one report per line whose count is greater than
  one as [`Duplicate report.count`](data-model.md) and
  [`Duplicate report.line text`](data-model.md) in `<count><tab><line-text>\n`
  form to an `io.Writer`. Return read/write errors to the caller. Re-run
  `go test ./ch1/dup` and confirm all cases from T001 pass.

**Checkpoint**: US1 behavior is verified by package unit tests without needing
the CLI entry point.

## Phase 4: User Story 2 - Run the example as a package (Priority: P1)

**Goal**: Make `ch1/dup` runnable from the repository root with stdin → stdout wiring.

**Independent test**: From the repository root, run
`printf 'a\nb\na\n' | go run ./ch1/dup` and confirm a successful exit with the
single report `2\ta` per [the CLI contract](contracts/cli.md).

- [X] T003 [US2] Rewrite `ch1/dup/main.go` as a thin entry point that calls the
  `ch1/dup/dup.go` API with `os.Stdin` and `os.Stdout`, exits successfully on
  finite input, and reports failures with a nonzero exit. Keep `main` as its
  only definition. Remove any duplicated counting/reporting logic so the package
  matches [the CLI contract](contracts/cli.md) and remains runnable with
  `go run ./ch1/dup`.

**Checkpoint**: US1 tests still pass and the CLI is runnable for US2.

## Phase 5: Polish & Cross-Cutting Concerns

- [X] T004 Format `ch1/dup/main.go`, `ch1/dup/dup.go`, and `ch1/dup/dup_test.go`
  with `gofmt`, confirm `gofmt -l ch1/dup` prints nothing, and run all unit
  tests with `go test -v ./ch1/dup`. Confirm the expected cases ran;
  `[no test files]` is not a pass.
- [X] T005 Run the CLI cases in `specs/002-dup-line-counts/quickstart.md` against
  [the CLI contract](contracts/cli.md), checking report sets and exit status
  (including the multi-duplicate order-independent case).

## Dependencies & Execution Order

### Phase Dependencies

- **Setup / Foundational**: Nothing to do; stories can start immediately.
- **User Story 1 (Phase 3)**: No infrastructure dependencies.
- **User Story 2 (Phase 4)**: Depends on US1's `ch1/dup/dup.go` API (T002).
- **Polish (Phase 5)**: Depends on US1 and US2 completion.

### User Story Dependencies

- **User Story 1 (P1)**: Independent; deliverable via unit tests alone.
- **User Story 2 (P1)**: Depends on T002; independently testable via
  `go run ./ch1/dup` once wired.

### Within Each User Story

- US1: T001 (failing tests) → T002 (implementation until tests pass).
- US2: T003 after T002.
- Polish: T004 → T005.

### Parallel Opportunities

- T001 is marked [P] relative to other files, but T002 must follow once the
  failing tests exist; do not parallelize T001 with T002.
- No other parallel tasks: US2 and polish share the same package files.

## Parallel Example: User Story 1

```bash
# Not applicable beyond writing T001 first, then implementing T002 in ch1/dup/dup.go.
# US1 has a single test file and a single implementation file in sequence.
```

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete T001–T002 so duplicate detection is covered by `go test ./ch1/dup`.
2. **STOP and VALIDATE**: Confirm VR-001 cases pass.
3. Continue with US2 for the runnable CLI.

### Incremental Delivery

1. US1 → package tests green (MVP for counting/reporting).
2. US2 → `go run ./ch1/dup` matches the CLI contract.
3. Polish → formatting, verbose tests, and quickstart CLI checks.

## Notes

- [P] tasks = different files, no dependencies on incomplete work.
- [Story] labels map to US1/US2 from [spec.md](spec.md).
- All constrained fields from [data-model.md](data-model.md) are cited in T001/T002.
- Commit after each task or logical group.
- Avoid leaving counting logic only in `main.go`; that fails the constitution
  unit-test gate for non-exempt source files.

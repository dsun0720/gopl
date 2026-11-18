# Tasks: Echo Arguments with Loops

**Input**: [Implementation plan](plan.md), [specification](spec.md), [research](research.md),
[data model](data-model.md), [CLI contract](contracts/cli.md), and [quickstart](quickstart.md).

**Tests**: Add exact-output unit tests for `ch1/echo/loops.go` per
[VR-001](spec.md#verification-requirements). `main.go` remains an entry-point-only
file and does not need dedicated unit tests.

## Phase 1: Setup

No tasks. Reuse the Go module declared in `go.mod` and the existing `ch1/echo` package.

## Phase 2: Foundational

No tasks. This standalone command needs no shared infrastructure or dependencies.

## Phase 3: User Story 1 - Compare equivalent loop forms (Priority: P1)

**Goal**: Show four distinct loop forms independently echoing the same arguments.

**Independent test**: Verify all 16 loop/input combinations with `go test ./ch1/echo`.
Run the four CLI cases in [quickstart.md](quickstart.md) to verify exact stdout,
no extra output, and successful exit. Inspect the source to confirm each loop
processes the arguments independently.

- [X] T001 [US1] Implement the counted, condition-only, `range`, and unconditional loop with explicit termination in `ch1/echo/loops.go`. Have each loop independently process [`Invocation arguments.values`](data-model.md) (the ordered list in the **Invocation arguments** row) and produce [`Echo result.text`](data-model.md) (the newline-terminated copy in the **Echo result** row). Use argument position, not accumulated text, to decide where to add separators.
- [X] T002 [US1] In `ch1/echo/main.go`, read the [`Invocation arguments.values`](data-model.md) list once, call all four demonstrations with that list, and remove the non-loop demonstration, separator lines, and `END` output so the command matches [the CLI contract](contracts/cli.md).
- [X] T003 [US1] Add table-driven tests in `ch1/echo/loops_test.go` for the four loop functions using the planned `io.Writer` interface and a fresh buffer per case. Assert exact output, including the final newline, for `hello world`, one argument containing `Go  loops`, an empty first argument followed by `word`, and no arguments. Run `go test ./ch1/echo` before the refactor and confirm that the new tests cannot pass against the current stdout-only functions.
- [X] T004 [US1] Update all four demonstrations in `ch1/echo/loops.go` to accept an `io.Writer` and return write errors, preserving the behavior and model references in T001. Update `ch1/echo/main.go` to pass `os.Stdout` and report write failures with a nonzero exit; keep `main` as its only definition. Re-run `go test ./ch1/echo` and confirm all 16 cases pass.

**Checkpoint**: US1 is runnable with `go run ./ch1/echo hello world`.

## Phase 4: Polish & Cross-Cutting Concerns

- [X] T005 Format `ch1/echo/main.go`, `ch1/echo/loops.go`, and `ch1/echo/loops_test.go` with `gofmt`, confirm `gofmt -l ch1/echo` prints nothing, and run all unit tests with `go test -v ./ch1/echo`. Confirm the expected cases ran; `[no test files]` is not a pass.
- [X] T006 Run the four CLI cases in `specs/001-hello-world-loops/quickstart.md` against the [CLI contract](contracts/cli.md), checking exact stdout and exit status. Inspect `ch1/echo/loops.go` for four independent loop implementations.

## Dependencies & Execution Order

`T001 → T002 → T003 → T004 → T005 → T006`. US1 has no other story or infrastructure
dependencies. T001 and T002 were completed before the test requirement was added.
There are no parallel tasks: add the tests before the writer refactor, then
validate the formatted package.

## Implementation Strategy

US1 is the entire MVP. Continue from T003, add test coverage without changing
CLI behavior, then run package-scoped and CLI verification.

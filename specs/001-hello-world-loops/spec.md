# Feature Specification: Echo Arguments with Loops

**Feature Branch**: Not created

**Created**: 2026-11-18

**Status**: Draft

**Input**: Demonstrate different Go `for` forms in the runnable `ch1/echo` command, with `ch1/echo/main.go` as the entry point and each form echoing the same arguments supplied at launch. This is a learning variant of [Chapter 1.1: Hello, World](https://golang-china.github.io/gopl-zh/ch1/ch1-01.html), not a reproduction of the book's example.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Compare equivalent loop forms (Priority: P1)

As a Go learner, I want to supply my own command-line arguments and see four loop forms echo the same values, so I can compare their behavior with different inputs.

**Why this priority**: Seeing the four forms side by side is the purpose of the example.

**Independent Test**: Run the example with different arguments and verify that each run echoes its own input four times. Inspect the example to identify the four distinct loop forms.

**Acceptance Scenarios**:

1. **Given** a learner at the repository root, **When** they run the `ch1/echo` command with arguments `hello` and `world`, **Then** the program exits successfully and prints exactly four lines of `hello world` to standard output, with no other output.
2. **Given** a learner supplies one quoted argument containing `Go  loops`, **When** they run the example, **Then** each of the four outputs is `Go  loops`, preserving the two spaces.
3. **Given** no arguments, **When** the learner runs the example, **Then** it exits successfully after printing four blank lines without prompting for input.
4. **Given** the example source, **When** the learner reviews the four demonstrations, **Then** they can identify one counted loop, one condition-only loop, one `range` loop, and one unconditioned loop with explicit termination, each producing its own copy of the output.

### Edge Cases

- No arguments means empty input, not a default greeting.
- Quoted whitespace and empty arguments are preserved; input is not trimmed or split again.
- The unconditioned loop must terminate without waiting for interactive input.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `ch1/echo` command MUST be runnable as a package with `ch1/echo/main.go` as its entry point. It MUST take its input from arguments supplied at launch, excluding the program name, and MUST NOT read interactive input.
- **FR-002**: Each of four distinct Go `for` forms MUST independently process the same argument list: a counted loop, a condition-only loop, a `range` loop, and an unconditioned loop with explicit termination.
- **FR-003**: Each demonstration MUST echo all arguments in order, preserving their contents and putting one space between adjacent arguments, followed by one newline. A successful run MUST exit successfully with no other output.
- **FR-004**: With no arguments, each demonstration MUST print only its terminating newline.

### Verification Requirements

- **VR-001**: Per Core Principle IV of the [constitution](../../.specify/memory/constitution.md), automated unit tests MUST verify the exact result of each loop demonstration for multiple arguments in order, quoted whitespace, an empty argument, and no arguments. The entry-point file is exempt from dedicated unit tests while `main` is its only definition.
- **VR-002**: All unit tests in the affected `ch1/echo` package MUST pass. A package run reporting no test files does not satisfy this requirement, and unrelated packages are not part of this acceptance gate.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: For each of two runs with different non-empty argument lists, standard output contains exactly four copies of that run's arguments joined by spaces, with one terminating newline per copy.
- **SC-002**: A run with no arguments exits successfully and produces exactly four newline characters.
- **SC-003**: All four demonstrations pass the four required input cases in automated unit tests.

## Assumptions

- “Different uses” means the four loop forms already illustrated in the chapter's `echo` example, not every possible use of loops.
- `hello world` is example input, not a built-in value. The existing separator and `END` output are outside the requested demonstration.
- The learner already has the Go toolchain required to run the repository's examples; other examples are out of scope.

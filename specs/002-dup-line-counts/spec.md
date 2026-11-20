# Feature Specification: Duplicate Line Counts

**Feature Branch**: Not created

**Created**: 2026-11-19

**Status**: Draft

**Input**: User description: "Create a module under ch1/dup. It must read from stdin and write duplicate lines plus their occurrence counts to stdout. The module must be runnable."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Find duplicate lines from standard input (Priority: P1)

As a Go learner, I want to pipe or type lines into the `ch1/dup` command and see which lines appear more than once along with how many times each appeared, so I can practice reading standard input and reporting duplicate content.

**Why this priority**: Detecting and reporting duplicate lines is the entire purpose of this example.

**Independent Test**: Feed known input with some repeated lines and verify that only those duplicates appear on standard output with the correct counts.

**Acceptance Scenarios**:

1. **Given** a learner at the repository root, **When** they run the `ch1/dup` command and supply standard input containing two identical non-empty lines and one unique line, then end input, **Then** the program exits successfully and prints exactly one result line showing count `2` and that duplicated text.
2. **Given** standard input where one line appears three times and another appears twice, **When** the learner runs the command to completion, **Then** standard output reports both lines with counts `3` and `2` respectively, and reports nothing for lines that appeared only once.
3. **Given** standard input with no duplicated lines (each line unique, or empty input), **When** the learner runs the command to completion, **Then** the program exits successfully and produces no output lines.

---

### User Story 2 - Run the example as a package (Priority: P1)

As a Go learner, I want `ch1/dup` to be a runnable package with a clear entry point, so I can invoke it the same way as other chapter examples without extra setup steps.

**Why this priority**: The request explicitly requires a runnable module under `ch1/dup`.

**Independent Test**: From the repository root, run the `ch1/dup` package with sample input and confirm it starts, reads input, and exits successfully.

**Acceptance Scenarios**:

1. **Given** the repository root and a Go toolchain, **When** the learner runs the `ch1/dup` package with any finite standard input, **Then** the command starts from `ch1/dup/main.go` as its entry point and exits successfully after processing that input.

---

### Edge Cases

- Empty standard input produces no duplicate report lines and still exits successfully.
- Lines that appear exactly once are omitted from the report.
- Leading or trailing spaces on a line are part of the line identity; two lines that differ only by whitespace are not treated as the same duplicate.
- An empty line that appears more than once is reported like any other duplicated line.
- Input ends when the stream reaches end-of-file; the command does not wait indefinitely after EOF.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `ch1/dup` command MUST be runnable as a package with `ch1/dup/main.go` as its entry point.
- **FR-002**: The command MUST read its input exclusively from standard input as a sequence of lines until end-of-file.
- **FR-003**: After reading completes, the command MUST write to standard output every line whose occurrence count is greater than one, together with that occurrence count.
- **FR-004**: Each reported result MUST appear on its own output line in the form `<count><tab><line-text>`, where `<count>` is the number of times that line occurred and `<line-text>` is the exact line content without its terminating newline.
- **FR-005**: Lines that occur exactly once MUST NOT appear in the output.
- **FR-006**: A successful run with finite input MUST exit successfully. When there are no duplicates, standard output MUST contain no result lines.

### Verification Requirements

- **VR-001**: Per Core Principle IV of the [constitution](../../.specify/memory/constitution.md), automated unit tests MUST verify counting and reporting behavior for inputs with duplicates, unique-only lines, empty input, empty duplicated lines, and whitespace-sensitive line identity. The entry-point file is exempt from dedicated unit tests while `main` is its only definition.
- **VR-002**: All unit tests in the affected `ch1/dup` package MUST pass. A package run reporting no test files does not satisfy this requirement, and unrelated packages are not part of this acceptance gate.

### Key Entities

- **Input Line**: One unit of text read from standard input, delimited by newlines; identity includes exact character content.
- **Duplicate Report**: A pairing of an input line with an occurrence count greater than one, emitted once per such line.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: For a prepared input with known duplicates, 100% of lines that appear more than once are reported with the correct count, and 0% of unique-only lines appear in the output.
- **SC-002**: A learner can run `ch1/dup` from the repository root with piped or typed input and complete a duplicate check in under one minute for inputs of a few dozen lines.
- **SC-003**: Automated unit tests covering the required duplicate-detection cases all pass for the `ch1/dup` package.

## Assumptions

- “Duplicate lines and their occurrence counts” means report only lines that appear more than once, matching the Chapter 1 duplicate-finding example, not a full frequency listing of every line.
- Output order among different duplicate lines is unspecified; correctness is judged by the set of reported line/count pairs.
- This feature covers reading standard input only; reading named files from arguments is out of scope unless added later as a separate exercise.
- The learner already has the Go toolchain required to run the repository's examples; other chapter examples are out of scope for this feature.
- Existing scaffolding under `ch1/dup` may be reused or adjusted so the package remains runnable and meets these requirements.

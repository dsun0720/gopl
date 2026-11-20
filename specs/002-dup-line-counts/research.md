# Research: Duplicate Line Counts

## Input source

- **Decision**: Read lines only from standard input until end-of-file.
- **Rationale**: The feature description and assumptions limit scope to stdin;
  file-path arguments belong to later chapter variants.
- **Alternatives considered**: Accepting named files would expand beyond the
  stated scope and mix exercise stages. Mixing args and stdin would blur the
  learning focus of this first form.

## Duplicate reporting rule

- **Decision**: Emit a report only for lines whose occurrence count is greater
  than one, in the form `<count><tab><line-text>` followed by a newline.
- **Rationale**: Matches the Chapter 1 duplicate-finding example and the
  learner-facing requirement to show repeated lines with their counts.
- **Alternatives considered**: Printing every line's frequency would change the
  meaning of "duplicate." Omitting the count would fail the stated output need.

## Line identity

- **Decision**: Treat each newline-delimited record's exact text (without the
  terminating newline) as the identity key, including empty strings and
  leading/trailing spaces.
- **Rationale**: Exact matching is the straightforward model for this exercise
  and matches the acceptance edge cases.
- **Alternatives considered**: Trimming or case-folding would hide differences
  the learner typed and is not requested.

## Output order

- **Decision**: Leave the order of distinct duplicate reports unspecified;
  validate the set of line/count pairs.
- **Rationale**: The chapter example accumulates counts in a map and iterates
  it; requiring a stable order would add constraints unrelated to the learning
  goal.
- **Alternatives considered**: Sorting by count or first-seen order would be
  clearer for humans but is not required by the spec.

## Testable package layout

- **Decision**: Move counting and reporting into `dup.go` with `io.Reader` /
  `io.Writer` parameters; keep `main.go` as a thin entry point. Use
  table-driven unit tests with buffers, asserting report sets rather than
  output byte order when multiple duplicates exist.
- **Rationale**: Constitution Principle IV requires tests for non-exempt source
  files. Injected streams avoid mutating process-wide stdin/stdout while still
  covering the behavior the CLI uses.
- **Alternatives considered**: Leaving all logic in `main.go` would evade unit
  tests but fail the constitution gate. Subprocess-only tests verify the CLI
  but are slower and harder to assert unordered multi-line reports cleanly.

# Research: Echo Arguments with Loops

## Input boundary

- **Decision**: Use only arguments supplied at launch, excluding the program name.
- **Rationale**: This preserves argument order and the exact strings supplied by
  the caller, including empty arguments and quoted whitespace.
- **Alternatives considered**: Reading standard input would introduce an
  interactive step; splitting or trimming would alter the supplied values.

## Independent loop demonstrations

- **Decision**: Retain the counted, condition-only, `range`, and unconditional
  loops. Add separators by argument position and terminate the unconditional
  loop when the argument list is exhausted.
- **Rationale**: All four loops perform the same operation. A positional rule
  preserves separators next to empty arguments and works for zero arguments.
- **Alternatives considered**: Joining once and printing the result four times
  would not demonstrate four ways to process the input. Testing whether the
  accumulated text is empty would drop a separator after an empty first argument.

## Exact-output unit tests

- **Decision**: Have each loop write to a supplied `io.Writer`. Use table-driven
  tests with a fresh `bytes.Buffer` for each loop and case.
- **Rationale**: This checks the output of each loop, including its terminating
  newline, without redirecting process-wide stdout. The CLI continues to write
  to standard output.
- **Alternatives considered**: Subprocess tests verify the whole CLI but do not
  isolate each loop. Replacing `os.Stdout` in tests mutates process-wide state.

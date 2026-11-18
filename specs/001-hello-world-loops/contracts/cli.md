# CLI Contract: Echo Arguments with Loops

**Invocation**: `go run ./ch1/echo [argument ...]` from the repository root.

The command accepts zero or more positional arguments as supplied by the shell.
It does not read standard input. Each of the four loop demonstrations writes one
copy of the text defined in the [data model](../data-model.md) to standard output.
A successful invocation exits with status zero and prints no headers, separators,
or footer beyond the spaces and terminating newlines defined by that model.

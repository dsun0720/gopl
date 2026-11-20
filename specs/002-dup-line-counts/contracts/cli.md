# CLI Contract: Duplicate Line Counts

**Invocation**: `go run ./ch1/dup` from the repository root, with line-oriented
input supplied on standard input until end-of-file.

The command accepts no required positional arguments for this feature. It reads
standard input only and does not open named files. After EOF, it writes zero or
more duplicate reports to standard output as defined in the
[data model](../data-model.md): each report is one line of
`<count><tab><line-text>`, emitted only for lines whose count is greater than
one. Report order is unspecified. A successful invocation exits with status
zero and prints no headers, prompts, or other text beyond those reports.

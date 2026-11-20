# Data Model: Duplicate Line Counts

| Entity | Fields and rules |
|--------|------------------|
| Input line | One newline-delimited text unit from standard input. Identity is the exact character content without the terminating newline, including empty content and surrounding spaces. |
| Occurrence count | Non-negative integer for how many times a given input line appears in one run. Starts at zero; each matching line increments by one. |
| Duplicate report | Emitted only when occurrence count is greater than one. Fields: count (integer > 1) and line text (exact input-line content). One report per distinct duplicated line. |

Nothing is persisted across runs. When no line has count greater than one, the
report set is empty. Order among reports is unspecified.

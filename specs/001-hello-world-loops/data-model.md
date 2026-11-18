# Data Model: Echo Arguments with Loops

| Entity | Fields and rules |
|--------|------------------|
| Invocation arguments | Ordered list of zero or more strings. Excludes the program name; each value is kept unchanged, including empty strings. |
| Echo result | One per loop form. Contains all argument values in order, separated by one space, followed by one newline. With no arguments it contains only the newline. |

All four results derive independently from the same argument list. Nothing is persisted.

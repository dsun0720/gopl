# Quickstart: Validate Duplicate Line Counts

From the repository root, with Go 1.26 or newer, run these checks after implementation.
See the [CLI contract](contracts/cli.md) for the expected behavior and the
[data model](data-model.md) for line identity and report rules.

```sh
gofmt -l ch1/dup
go test -v ./ch1/dup
```

`gofmt -l` must print nothing. The test run must show passing cases for
duplicates, unique-only lines, empty input, empty duplicated lines, and
whitespace-sensitive identity. A result of `[no test files]` does not pass.

| Run | Expected standard output |
|-----|--------------------------|
| `printf 'a\nb\na\n' \| go run ./ch1/dup` | Exactly one line: `2` then a tab then `a` |
| `printf 'x\ny\nx\ny\nx\n' \| go run ./ch1/dup` | Two reports: `3` tab `x`, and `2` tab `y` (either order) |
| `printf 'only\nonce\n' \| go run ./ch1/dup` | No output |
| `printf '' \| go run ./ch1/dup` | No output |
| `printf '\n\n' \| go run ./ch1/dup` | Exactly one line: `2` then a tab (empty line text) |
| `printf ' hi\nhi\n hi\n' \| go run ./ch1/dup` | Exactly one line: `2` then a tab then ` hi` (`hi` alone is not reported) |

Every invocation must exit successfully after EOF without printing headers or
extra text beyond the reports defined by the contract.

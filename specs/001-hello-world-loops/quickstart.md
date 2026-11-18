# Quickstart: Validate Echo Arguments with Loops

From the repository root, with Go 1.26 or newer, run these checks after implementation.
See the [CLI contract](contracts/cli.md) for the expected behavior.

```sh
gofmt -l ch1/echo
go test -v ./ch1/echo
```

`gofmt -l` must print nothing. The test run must show passing cases for each
of the four loop forms with multiple arguments, quoted whitespace, an empty
argument, and no arguments. A result of `[no test files]` does not pass.

| Run | Expected standard output |
|-----|--------------------------|
| `go run ./ch1/echo hello world` | Four `hello world` lines |
| `go run ./ch1/echo 'Go  loops'` | Four `Go  loops` lines, each preserving two spaces |
| `go run ./ch1/echo '' word` | Four ` word` lines, each starting with one space |
| `go run ./ch1/echo` | Exactly four newline characters |

Every invocation must exit successfully without prompting or printing extra text.
For the no-argument case, `go run ./ch1/echo | od -An -t x1` should show
`0a 0a 0a 0a`. Inspect `ch1/echo/loops.go` to confirm that the four distinct loop
forms each process the argument list rather than reprinting one shared result.

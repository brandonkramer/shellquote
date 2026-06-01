# shellquote

Shell argument quoting and prompt-file helpers for command templates.

## Install

From [pkg.go.dev](https://pkg.go.dev/github.com/brandonkramer/shellquote):

```bash
go get github.com/brandonkramer/shellquote@v0.1.0
```

## Usage

```go
path, err := shellquote.WritePrompt(runDir, promptContent)
cmd := shellquote.SubstitutePrompt("cat {prompt}", path)
quoted := shellquote.Quote("/tmp/a file")
```

Constants:

- `shellquote.DefaultPromptFile` — `"prompt.md"`
- `shellquote.PromptPlaceholder` — `"{prompt}"`

## Development

Lefthook and golangci-lint are pinned in `go.mod` as **tools** (dev-only). Install git hooks once per clone:

```bash
make install-hooks
```

```bash
make check
make test
make lint
```

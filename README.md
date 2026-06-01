# shellquote

Shell argument quoting and prompt-file helpers for command templates.

## Install

```bash
go get github.com/brandonkramer/shellquote
```

## Usage

```go
cmd := shellquote.SubstitutePrompt("cat {prompt}", "/tmp/run/prompt.md")
path, err := shellquote.WritePrompt(runDir, promptContent)
```

package shellquote_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brandonkramer/shellquote"
)

func TestQuote(t *testing.T) {
	t.Parallel()

	if got := shellquote.Quote("a'b"); got != "'a'\\''b'" {
		t.Fatalf("quote=%q", got)
	}
}

func TestSubstitutePrompt(t *testing.T) {
	t.Parallel()

	if got := shellquote.SubstitutePrompt("echo hi", "/p"); got != "echo hi" {
		t.Fatalf("got=%q", got)
	}
	got := shellquote.SubstitutePrompt("cat {prompt}", "/tmp/prompt.md")
	if got == "cat {prompt}" {
		t.Fatalf("expected substitution, got %q", got)
	}
	if !strings.Contains(got, shellquote.Quote("/tmp/prompt.md")) {
		t.Fatalf("expected quoted path in %q", got)
	}
}

func TestPromptPath(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	got := shellquote.PromptPath(dir)
	want := filepath.Join(dir, shellquote.DefaultPromptFile)
	if got != want {
		t.Fatalf("path=%q want %q", got, want)
	}
}

func TestWritePrompt(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path, err := shellquote.WritePrompt(dir, "hello")
	if err != nil {
		t.Fatal(err)
	}
	if path != shellquote.PromptPath(dir) {
		t.Fatalf("path=%q", path)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "hello" {
		t.Fatalf("data=%q err=%v", data, err)
	}
}

func TestWritePromptError(t *testing.T) {
	t.Parallel()

	_, err := shellquote.WritePrompt("/", "x")
	if err == nil {
		t.Fatal("expected write error")
	}
	if !strings.Contains(err.Error(), "shellquote: write prompt:") {
		t.Fatalf("err=%v", err)
	}
}

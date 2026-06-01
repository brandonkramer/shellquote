package shellquote_test

import (
	"os"
	"path/filepath"
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
}

func TestWritePrompt(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path, err := shellquote.WritePrompt(dir, "hello")
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(dir, shellquote.DefaultPromptFile)
	if path != want {
		t.Fatalf("path=%q want %q", path, want)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "hello" {
		t.Fatalf("data=%q err=%v", data, err)
	}
}

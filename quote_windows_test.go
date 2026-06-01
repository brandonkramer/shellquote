//go:build windows

package shellquote_test

import (
	"testing"

	"github.com/brandonkramer/shellquote"
)

func TestQuoteWindows(t *testing.T) {
	t.Parallel()

	if got := shellquote.Quote(""); got != `""` {
		t.Fatalf("empty=%q", got)
	}
	if got := shellquote.Quote(`a"b`); got != `"a\"b"` {
		t.Fatalf("quoted=%q", got)
	}
	if got := shellquote.Quote("plain"); got != "plain" {
		t.Fatalf("plain=%q", got)
	}
}

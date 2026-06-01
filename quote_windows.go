//go:build windows

package shellquote

import "strings"

// Quote wraps s for Windows cmd-style argument passing.
func Quote(s string) string {
	if s == "" {
		return `""`
	}
	if !strings.ContainsAny(s, " \t\n\v\f\"\\") {
		return s
	}
	return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
}

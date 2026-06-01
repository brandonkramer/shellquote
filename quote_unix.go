//go:build unix

package shellquote

import "strings"

// Quote wraps s in single quotes for POSIX shells.
func Quote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

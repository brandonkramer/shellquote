package shellquote

import (
	"os"
	"path/filepath"
)

// PromptPath returns the default prompt file path under runDir.
func PromptPath(runDir string) string {
	return filepath.Join(runDir, DefaultPromptFile)
}

// WritePrompt writes content to the default prompt file under runDir.
func WritePrompt(runDir, content string) (string, error) {
	path := PromptPath(runDir)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		return "", err
	}
	return path, nil
}

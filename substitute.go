package shellquote

import "strings"

const (
	// DefaultPromptFile is the default prompt filename under a run directory.
	DefaultPromptFile = "prompt.md"
	// PromptPlaceholder is substituted with a quoted prompt path in templates.
	PromptPlaceholder = "{prompt}"
)

// SubstitutePrompt replaces PromptPlaceholder with a quoted promptPath in template.
func SubstitutePrompt(template, promptPath string) string {
	if !strings.Contains(template, PromptPlaceholder) {
		return template
	}
	return strings.ReplaceAll(template, PromptPlaceholder, Quote(promptPath))
}

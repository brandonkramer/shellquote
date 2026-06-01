// Package shellquote provides shell argument quoting and prompt-file helpers.
//
// Use [Quote] and [SubstitutePrompt] when building shell command templates,
// and [WritePrompt] to materialize prompt content under a run directory.
//
// Example:
//
//	path, err := shellquote.WritePrompt(runDir, prompt)
//	cmd := shellquote.SubstitutePrompt("cat {prompt}", path)
package shellquote

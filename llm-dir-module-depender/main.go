// A generated module for LlmDirModuleDepender functions
// this module depends on an LLM-using module, but doesn't touch the LLM API by itself
package main

import (
	"context"
	"dagger/llm-dir-module-depender/internal/dagger"
)

type LlmDirModuleDepender struct {
	Model string // +private
}

func New(
	model string, // +optional
) *LlmDirModuleDepender {
	return &LlmDirModuleDepender{model}
}

// Returns a container that echoes whatever string argument is provided
func (m *LlmDirModuleDepender) PromptDep(ctx context.Context, stringArg string) (string, error) {
	return dag.LlmTestModule(dagger.LlmTestModuleOpts{Model: m.Model}).Prompt(ctx, stringArg)
}

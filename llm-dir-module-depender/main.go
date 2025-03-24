// A generated module for LlmDirModuleDepender functions
// this module depends on an LLM-using module, but doesn't touch the LLM API by itself
package main

import (
	"context"
	"dagger/llm-dir-module-depender/internal/dagger"
)

type LLMDirModuleDepender struct {
	Model string // +private
}

func New(
	model string, // +optional
) *LLMDirModuleDepender {
	return &LLMDirModuleDepender{model}
}

// Returns a container that echoes whatever string argument is provided
func (m *LLMDirModuleDepender) PromptDep(ctx context.Context, stringArg string) (string, error) {
	return dag.LLMTestModule(dagger.LLMTestModuleOpts{Model: m.Model}).Prompt(ctx, stringArg)
}

func (m *LLMDirModuleDepender) Save(ctx context.Context, stringArg string) (string, error) {
	return dag.LLMTestModule(dagger.LLMTestModuleOpts{Model: m.Model}).Save(ctx, stringArg)
}

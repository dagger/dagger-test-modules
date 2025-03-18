// A generated module for Llm functions
// This module is necessary for testing allow-llm policy enforcement around git-sourced modules and is called by its parent LlmDirModuleDepender
package main

import (
	"context"
	"dagger/llm/internal/dagger"
)

type LlmTestModule struct {
	Model string // +private
}

func New(
	model string, // +optional
) *LlmTestModule {
	return &LlmTestModule{model}
}

// access the LLM api in the most minimal way: prompt it, then retrieve history
func (m *LlmTestModule) Prompt(ctx context.Context, stringArg string) (string, error) {
	return dag.Llm(dagger.LlmOpts{Model: m.Model}).WithPrompt(stringArg).LastReply(ctx)
}

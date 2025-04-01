// A generated module for LLM functions
// This module is necessary for testing allow-llm policy enforcement around git-sourced modules and is called by its parent LLMDirModuleDepender
package main

import (
	"context"
	"dagger/llm/internal/dagger"

	"github.com/google/uuid"
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
	return m.llm(stringArg).LastReply(ctx)
}

// this is a hack until we can return the LLM type and replay history
func (m *LlmTestModule) Save(ctx context.Context, stringArg string) (string, error) {
	return m.llm(stringArg).HistoryJSON(ctx)
}

func (m LlmTestModule) llm(stringArg string) *dagger.LLM {
	return dag.
		LLM(dagger.LLMOpts{Model: m.Model}).
		WithEnv(dag.Env().WithStringInput("CACHE_BUSTER", uuid.NewString(), "a random string only meant for busting caches, not actual use.")).
		WithPrompt(stringArg)
}

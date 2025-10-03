package main

import (
	"context"
	"encoding/json"

	"main/internal/dagger"
)

type CoolSdk struct{}

func (m *CoolSdk) ModuleDefs(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File, outputFilePath string) (*dagger.Container, error) {
	mod := modSource.WithSDK("go").AsModule()
	modID, err := mod.ID(ctx)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(modID)
	if err != nil {
		return nil, err
	}
	return dag.Container().
		From("alpine").
		WithNewFile(outputFilePath, string(b)).
		WithEntrypoint([]string{
			"sh", "-c", "",
		}), nil
}

func (m *CoolSdk) ModuleRuntime(modSource *dagger.ModuleSource, introspectionJson *dagger.File) *dagger.Container {
	return modSource.WithSDK("go").AsModule().Runtime().WithEnvVariable("COOL", "true")
}

func (m *CoolSdk) Codegen(modSource *dagger.ModuleSource, introspectionJson *dagger.File) *dagger.GeneratedCode {
	return dag.GeneratedCode(modSource.WithSDK("go").AsModule().GeneratedContextDirectory())
}

func (m *CoolSdk) RequiredPaths() []string {
	return []string{
		"**/go.mod",
		"**/go.sum",
		"**/go.work",
		"**/go.work.sum",
		"**/vendor/",
		"**/*.go",
	}
}

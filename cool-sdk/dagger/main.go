package main

import "main/internal/dagger"

type CoolSdk struct{}

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

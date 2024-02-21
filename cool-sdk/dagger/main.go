package main

type CoolSdk struct{}

func (m *CoolSdk) ModuleRuntime(modSource *ModuleSource, introspectionJson string) *Container {
	return modSource.WithSDK("go").AsModule().Runtime().WithEnvVariable("COOL", "true")
}

func (m *CoolSdk) Codegen(modSource *ModuleSource, introspectionJson string) *GeneratedCode {
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

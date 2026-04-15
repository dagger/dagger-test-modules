// Package engine_dep is a minimal sibling module used by greeter as a
// relative dependency (`../engine-dep`). Its only purpose is to force the
// greeter loader to ALSO resolve another module inside the same remote
// repo, mirroring toolchains/java-sdk-dev's `dependencies: [{source:
// "../engine-dev"}]` pattern that is suspected to perturb the parent's
// context-directory resolution.
package main

type EngineDep struct{}

// Ping is a no-op function; it exists so the module has something to
// expose and the dep load is real.
func (m *EngineDep) Ping() string {
	return "pong"
}

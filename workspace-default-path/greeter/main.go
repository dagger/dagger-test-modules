// Package greeter reproduces a bug where a toolchain-style module loaded
// via a remote git ref cannot walk into a workspace subdirectory through a
// Directory argument annotated with defaultPath="/" + ignore patterns.
//
// When loaded as github.com/dagger/dagger-test-modules/workspace-default-path/greeter@<ref>,
// the module's defaultPath="/" must resolve to the repo root (where .git and
// the root dagger.json live). The ignore patterns exclude everything except
// workspace-default-path/target-subdir/, so workspace.Directory(
// "workspace-default-path/target-subdir") must return a directory whose
// contents include hello.txt.
//
// This mirrors toolchains/java-sdk-dev's production pattern:
//
//	pub workspace: Directory! @defaultPath(path: "/") @ignorePatterns(patterns: [
//	    "*",
//	    "!sdk/java/runtime/images/",
//	    ...
//	])
//
//	workspace.directory("sdk/java/runtime/images/maven")
package main

import (
	"context"

	"dagger/greeter/internal/dagger"
)

type Greeter struct{}

// Read returns the contents of target-subdir/hello.txt, reached by walking
// into the workspace root (defaultPath="/") and then into the target
// subdirectory. The ignore patterns must re-include that subdirectory for
// Read to find the file.
//
// It also calls engine-dep (a relative dependency) to force the parent's
// load path to actually resolve the sibling module — mirroring the
// java-sdk-dev case where `../engine-dev` is fetched as part of the
// toolchain load.
func (m *Greeter) Read(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=["*", "!workspace-default-path/target-subdir/"]
	workspace *dagger.Directory,
) (string, error) {
	if _, err := dag.EngineDep().Ping(ctx); err != nil {
		return "", err
	}
	return workspace.
		Directory("workspace-default-path/target-subdir").
		File("hello.txt").
		Contents(ctx)
}

// WorkspaceEntries lists the entries at the workspace root after ignore
// patterns are applied. It's a sanity check that distinguishes
// "workspace root doesn't load" from "walking into a subdirectory fails" —
// only the latter is the target bug.
func (m *Greeter) WorkspaceEntries(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=["*", "!workspace-default-path/target-subdir/"]
	workspace *dagger.Directory,
) ([]string, error) {
	return workspace.Entries(ctx)
}

// ReadCheck exposes the same subdirectory access as Read through a
// check-annotated function. The production java-sdk-dev regression
// surfaces via @check functions (`test`, `lint`), so this path must pass
// once the bug is fixed.
//
// +check
func (m *Greeter) ReadCheck(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=["*", "!workspace-default-path/target-subdir/"]
	workspace *dagger.Directory,
) error {
	if _, err := dag.EngineDep().Ping(ctx); err != nil {
		return err
	}
	_, err := workspace.
		Directory("workspace-default-path/target-subdir").
		File("hello.txt").
		Contents(ctx)
	return err
}

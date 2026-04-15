// Package greeter reproduces a bug where a toolchain-style module loaded
// via a remote git ref cannot walk into a workspace *sub-sub*-directory
// through a Directory argument annotated with defaultPath="/" + ignore
// patterns.
//
// When loaded as github.com/dagger/dagger-test-modules/workspace-default-path/greeter@<ref>,
// the module's defaultPath="/" must resolve to the repo root. The ignore
// patterns un-exclude `workspace-default-path/target-subdir/` (the parent)
// and the function then accesses
// `workspace-default-path/target-subdir/maven` — a sub-directory of the
// un-excluded parent that is itself never named in the pattern list.
// That's the exact shape of the java-sdk-dev production failure:
//
//	pub workspace: Directory! @defaultPath(path: "/") @ignorePatterns(patterns: [
//	    "*",
//	    "!sdk/java/runtime/images/",      // un-exclude a parent
//	    ...
//	])
//
//	workspace.directory("sdk/java/runtime/images/maven")   // access sub of parent
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
		Directory("workspace-default-path/target-subdir/maven").
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
		Directory("workspace-default-path/target-subdir/maven").
		File("hello.txt").
		Contents(ctx)
	return err
}

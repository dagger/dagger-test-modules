// Package greeter reproduces a bug where a toolchain-style module loaded
// via a remote git ref cannot walk into a workspace sub-sub-directory
// through a constructor-injected Directory field annotated with
// defaultPath="/" + ignore patterns.
//
// This mirrors toolchains/java-sdk-dev exactly: the workspace is a
// TYPE FIELD (set once at construction) with @defaultPath/@ignorePatterns
// — not a per-function argument. Functions then access
// m.Workspace.Directory(...) through the stored field, so the
// default-path resolution happens at constructor time and its result is
// reused across calls. The real failure path is:
//
//	dagger -m github.com/.../workspace-default-path/greeter@<sha> call read
//	=> New(workspace=<defaultPath="/"+ignore>)
//	=> m.Workspace.Directory("workspace-default-path/target-subdir/maven")...
package main

import (
	"context"

	"dagger/greeter/internal/dagger"
)

type Greeter struct {
	// Workspace is stored on the receiver so function calls reuse the
	// same resolved Directory instance — matching dang's `pub workspace`
	// field storage model in java-sdk-dev.
	Workspace *dagger.Directory
}

// New injects the workspace via defaultPath + ignore at construction time.
// Subsequent function calls on the returned Greeter operate on the stored
// m.Workspace, NOT a freshly-resolved one.
func New(
	// +defaultPath="/"
	// +ignore=["*", "!workspace-default-path/target-subdir/"]
	workspace *dagger.Directory,
) *Greeter {
	return &Greeter{Workspace: workspace}
}

// Read walks into a sub-sub-directory of the un-excluded parent. The
// "maven" child is never named in the ignore pattern list — it must be
// pulled in transitively by the parent's trailing-slash un-exclusion.
func (m *Greeter) Read(ctx context.Context) (string, error) {
	if _, err := dag.EngineDep().Ping(ctx); err != nil {
		return "", err
	}
	return m.Workspace.
		Directory("workspace-default-path/target-subdir/maven").
		File("hello.txt").
		Contents(ctx)
}

// WorkspaceEntries is a sanity check: workspace root must load.
func (m *Greeter) WorkspaceEntries(ctx context.Context) ([]string, error) {
	return m.Workspace.Entries(ctx)
}

// ReadCheck exposes the sub-sub-directory access through a check.
// The production java-sdk-dev regression surfaces through @check
// functions (`test`, `lint`), so this path must pass once the bug is
// fixed.
//
// +check
func (m *Greeter) ReadCheck(ctx context.Context) error {
	if _, err := dag.EngineDep().Ping(ctx); err != nil {
		return err
	}
	_, err := m.Workspace.
		Directory("workspace-default-path/target-subdir/maven").
		File("hello.txt").
		Contents(ctx)
	return err
}

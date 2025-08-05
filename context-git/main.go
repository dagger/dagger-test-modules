package main

import (
	"context"
	"dagger/context-git/internal/dagger"
)

type ContextGit struct{}

func (m *ContextGit) TestRepoLocal(
	ctx context.Context,
	// +defaultPath="./.git"
	git *dagger.GitRepository,
) (string, error) {
	return m.commitAndRef(ctx, git.Head())
}

func (m *ContextGit) TestRepoRemote(
	ctx context.Context,
	// +defaultPath="https://github.com/dagger/dagger.git"
	git *dagger.GitRepository,
) (string, error) {
	return m.commitAndRef(ctx, git.Tag("v0.18.2"))
}

func (m *ContextGit) TestRefLocal(
	ctx context.Context,
	// +defaultPath="./.git"
	git *dagger.GitRef,
) (string, error) {
	return m.commitAndRef(ctx, git)
}

func (m *ContextGit) TestRefRemote(
	ctx context.Context,
	// +defaultPath="https://github.com/dagger/dagger.git#v0.18.3"
	git *dagger.GitRef,
) (string, error) {
	return m.commitAndRef(ctx, git)
}

func (m *ContextGit) commitAndRef(ctx context.Context, ref *dagger.GitRef) (string, error) {
	commit, err := ref.Commit(ctx)
	if err != nil {
		return "", err
	}
	reference, err := ref.Ref(ctx)
	if err != nil {
		return "", err
	}
	return reference + "@" + commit, nil
}

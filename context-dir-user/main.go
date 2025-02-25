package main

import (
	"dagger/context-dir-user/internal/dagger"
)

type ContextDirUser struct{}

func (m *ContextDirUser) AbsolutePath() *dagger.Directory {
	return dag.Test().AbsolutePath()
}

func (m *ContextDirUser) AbsolutePathSubdir() *dagger.Directory {
	return dag.Test().AbsolutePathSubdir()
}

func (m *ContextDirUser) RelativePath() *dagger.Directory {
	return dag.Test().RelativePath()
}

func (m *ContextDirUser) RelativePathSubdir() *dagger.Directory {
	return dag.Test().RelativePathSubdir()
}

package main

import (
	"dagger/test/internal/dagger"
)

type Test struct{}

func (m *Test) AbsolutePath(
	// +defaultPath="/"
	dir *dagger.Directory,
) *dagger.Directory {
	return dir
}

func (m *Test) RelativePath(
	// +defaultPath="."
	dir *dagger.Directory,
) *dagger.Directory {
	return dir
}

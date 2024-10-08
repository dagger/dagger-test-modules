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

func (m *Test) AbsolutePathSubdir(
	// +defaultPath="/data"
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


func (m *Test) RelativePathSubdir(
	// +defaultPath="./data/subdir"
	dir *dagger.Directory,
) *dagger.Directory {
	return dir
}
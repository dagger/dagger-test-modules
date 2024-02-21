package main

import (
	"context"
)

type Dep struct{}

func (m *Dep) Fn(ctx context.Context) (string, error) {
	depHi, err := dag.Dep2().Fn(ctx)
	if err != nil {
		return "", err
	}
	return "hi from dep " + depHi, nil
}

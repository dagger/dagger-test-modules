package main

import (
	"context"
)

type TopLevel struct{}

func (m *TopLevel) Fn(ctx context.Context) (string, error) {
	depHi, err := dag.Dep().Fn(ctx)
	if err != nil {
		return "", err
	}
	return "hi from top level " + depHi, nil
}

package main

import (
	"context"
)

type RootMod struct{}

func (m *RootMod) Fn(ctx context.Context) (string, error) {
	depHi, err := dag.DepAlias().Fn(ctx)
	if err != nil {
		return "", err
	}
	return "hi from root " + depHi, nil
}

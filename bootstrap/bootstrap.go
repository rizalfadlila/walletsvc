package bootstrap

import (
	"github.com/julo/walletsvc/bootstrap/deps"
	"github.com/julo/walletsvc/bootstrap/initiator"
)

const (
	defaultGracefulTimeout = "15s"
)

type bootstrap struct {
	deps *deps.Deps
}

func setup() *bootstrap {
	dependencies := initiator.New().
		InitBasic().
		InitRepositories().
		InitService().
		SetupDeps()

	return &bootstrap{deps: dependencies}
}

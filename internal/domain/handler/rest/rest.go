package rest

import (
	"github.com/julo/walletsvc/bootstrap/deps"
	"github.com/julo/walletsvc/internal/domain/handler/rest/api"
)

type (
	module struct {
		deps *deps.Deps
		err  chan error
	}

	Handler interface {
		Run()
		Error() <-chan error
	}
)

func New(deps *deps.Deps) Handler {
	handler := &module{deps: deps}

	apiOpts := api.Opts{
		Router:  handler.deps.Basic.Router,
		Auth:    handler.deps.Basic.Auth,
		Service: handler.deps.Service,
	}

	api.New(apiOpts).Register()

	return handler
}

func (h *module) Run() {
	h.err <- h.deps.Basic.Router.Start()
}

func (h *module) Error() <-chan error {
	return h.err
}

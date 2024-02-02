package bootstrap

import "github.com/julo/walletsvc/internal/domain/handler/rest"

type Rest struct {
	*bootstrap
	handler rest.Handler
}

func NewRest() *Rest {
	return &Rest{
		bootstrap: setup(),
	}
}

func (r *Rest) RegisterHandler() *Rest {
	r.handler = rest.New(r.deps)
	return r
}

func (r *Rest) Serve() {
	go r.handler.Run()

	r.errorHandler(r.handler.Error())
}

package api

import (
	"github.com/julo/walletsvc/bootstrap/deps"
	"github.com/julo/walletsvc/internal/infrastructure/auth"
	"github.com/julo/walletsvc/internal/infrastructure/middleware"
	"github.com/julo/walletsvc/internal/infrastructure/router"
)

type (
	API struct {
		router  *router.MyRouter
		auth    auth.Auth
		service *deps.Service
	}

	Opts struct {
		Router  *router.MyRouter
		Auth    auth.Auth
		Service *deps.Service
	}
)

func New(o Opts) *API {
	return &API{
		router:  o.Router,
		auth:    o.Auth,
		service: o.Service,
	}
}

func (a *API) Register() {
	a.v1()
}

func (a *API) v1() {
	a.router.Group("/v1", func(r *router.MyRouter) {
		r.POST("/init", a.CreateWalletAccount)
		r.Group("/wallet", func(r *router.MyRouter) {
			r.Use(middleware.NewJwt(a.auth))
			r.POST("", a.EnableWallet)
			r.GET("", a.ViewWallet)
		})
	})
}

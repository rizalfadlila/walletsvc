package router

import (
	"github.com/gofiber/fiber/v2"
)

func (r *MyRouter) Group(group string, fn func(r *MyRouter), handler ...Handler) {
	handlers := make([]fiber.Handler, 0)
	for _, v := range handler {
		handlers = append(handlers, r.convertHandler(v))
	}

	g := r.router.Group(group, handlers...)

	fn(&MyRouter{router: g})
}

func (r *MyRouter) Use(args ...interface{}) {
	r.router.Use(args...)
}

func (r *MyRouter) GET(path string, handler Handler) {
	r.router.Get(path, r.convertHandler(handler))
}

func (r *MyRouter) POST(path string, handler Handler) {
	r.router.Post(path, r.convertHandler(handler))
}

func (r *MyRouter) PUT(path string, handler Handler) {
	r.router.Put(path, r.convertHandler(handler))
}

func (r *MyRouter) DELETE(path string, handler Handler) {
	r.router.Delete(path, r.convertHandler(handler))
}

func (r *MyRouter) PATCH(path string, handler Handler) {
	r.router.Patch(path, r.convertHandler(handler))
}

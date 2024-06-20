package core

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Mux *echo.Echo
}

func (r *Router) Pre(middleware ...echo.MiddlewareFunc) {
	r.Mux.Pre(middleware...)
}

func (r *Router) Use(middleware ...echo.MiddlewareFunc) {
	r.Mux.Use(middleware...)
}

func (r *Router) Any(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return r.Mux.Any(path, h, m...)
}

func (r *Router) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.GET(path, h, m...)
}

func (r *Router) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.POST(path, h, m...)
}

func (r *Router) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.PATCH(path, h, m...)
}

func (r *Router) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.PUT(path, h, m...)
}

func (r *Router) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.DELETE(path, h, m...)
}

func (r *Router) RouteNotFound(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.RouteNotFound(path, h, m...)
}

func (r *Router) Match(methods []string, path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return r.Mux.Match(methods, path, h, m...)
}

func setRequestContext(ctx context.Context, r *http.Request, key, val any) {
	req := r.WithContext(context.WithValue(ctx, key, val))
	*r = *req
}

func customContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Context{c}
		return next(cc)
	}
}

func checkHtmx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if c.Request().Header.Get("Hx-Request") == "true" {
			setRequestContext(ctx, c.Request(), "hx-request", true)
			c.Response().Header().Set("Cache-Control", "no-store, max-age=0")
		}
		return next(c)
	}
}

func NewRouter() *Router {
	e := echo.New()
	e.Use(customContext)
	e.Use(checkHtmx)

	return &Router{Mux: e}
}

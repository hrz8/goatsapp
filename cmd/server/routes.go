package server

import (
	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/internal/exception"
	"github.com/hrz8/goatsapp/internal/homepage"
	"github.com/hrz8/goatsapp/internal/static"
	"github.com/hrz8/gofx"
)

func RegisterRouters(homepage *homepage.Handler, static *static.Handler, exception *exception.Handler) *gofx.Router {
	router := gofx.NewRouter()

	// views
	router.GET("/", homepage.Index)
	router.GET("/assets/*", static.Serve(assets.FS))

	// api

	// fallback
	router.RouteNotFound("", exception.NotFound)
	router.RouteNotFound("/*", exception.NotFound)

	return router
}

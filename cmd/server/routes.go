package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/internal/exception"
	"github.com/hrz8/goatsapp/internal/homepage"
	"github.com/hrz8/goatsapp/internal/project"
	"github.com/hrz8/goatsapp/internal/static"
	"github.com/hrz8/gofx"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.Validator.Struct(i)
}

func RegisterRouters(
	homepage *homepage.Handler,
	static *static.Handler,
	exception *exception.Handler,
	project *project.Handler,
) *gofx.Router {
	router := gofx.NewRouter()

	// global setup
	router.Mux.Validator = &CustomValidator{validator.New()}

	// views
	router.GET("/", homepage.Index)
	router.GET("/assets/*", static.Serve(assets.FS))

	// htmx/ajax
	router.POST("/projects/new", project.CreateProjectForm)

	// fallback
	router.RouteNotFound("", exception.NotFound)
	router.RouteNotFound("/*", exception.NotFound)

	return router
}

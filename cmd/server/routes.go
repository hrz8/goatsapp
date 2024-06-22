package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/goatsapp/internal/exception"
	"github.com/hrz8/goatsapp/internal/homepage"
	Project "github.com/hrz8/goatsapp/internal/project"
	"github.com/hrz8/goatsapp/internal/static"
	"github.com/hrz8/goatsapp/pkg/middleware"
	"github.com/hrz8/gofx"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.Validator.Struct(i)
}

func RegisterRouters(
	repo *dbrepo.Queries,
	// modules
	homepage *homepage.Handler,
	static *static.Handler,
	exception *exception.Handler,
	project *Project.Handler,
) *gofx.Router {
	router := gofx.NewRouter()

	// global setup
	router.Mux.Validator = &CustomValidator{validator.New()}
	router.Use(middleware.AppIDCookie(repo))

	// views
	router.GET("/", homepage.Index)
	router.GET("/assets/*", static.Serve(assets.FS))

	// htmx/ajax
	router.POST("/projects/new", project.CreateProjectForm)
	router.GET("/projects/selector", project.ListProjectSelector)

	// fallback
	router.RouteNotFound("", exception.NotFound)
	router.RouteNotFound("/*", exception.NotFound)

	return router
}

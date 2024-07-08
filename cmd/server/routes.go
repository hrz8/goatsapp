package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/hrz8/goatsapp/assets"
	Device "github.com/hrz8/goatsapp/internal/device"
	Exception "github.com/hrz8/goatsapp/internal/exception"
	Homepage "github.com/hrz8/goatsapp/internal/homepage"
	Middleware "github.com/hrz8/goatsapp/internal/middleware"
	Project "github.com/hrz8/goatsapp/internal/project"
	Static "github.com/hrz8/goatsapp/internal/static"

	"github.com/hrz8/gofx"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.Validator.Struct(i)
}

func RegisterRouters(
	middleware *Middleware.Handler,
	homepage *Homepage.Handler,
	device *Device.Handler,
	static *Static.Handler,
	exception *Exception.Handler,
	project *Project.Handler,
) *gofx.Router {
	router := gofx.NewRouter()
	router.Mux.File("/favicon.ico", "assets/favicon/favicon.ico")

	// global setup
	router.Mux.Validator = &CustomValidator{validator.New()}
	router.Use(middleware.ProjectIDCookie)

	// views
	router.GET("/", homepage.Index)
	router.GET("/devices", device.Index)
	router.GET("/project-settings", project.Settings)
	router.GET("/assets/*", static.Serve(assets.FS))

	// htmx/ajax
	router.POST("/projects/new", project.CreateProjectForm)
	router.GET("/projects/selector", project.ListProjectSelector)
	router.POST("/projects/switch", project.SwitchProject)
	router.POST("/devices/new", device.CreateDeviceForm)
	router.GET("/devices/qr", device.RequestQR)

	// fallback
	router.RouteNotFound("", exception.NotFound)
	router.RouteNotFound("/*", exception.NotFound)

	return router
}

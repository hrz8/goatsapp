package device

import (
	"net/http"

	"github.com/hrz8/goatsapp/web/template/component"
	"github.com/hrz8/goatsapp/web/template/page"
	"github.com/hrz8/gofx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	log *gofx.Logger
	svc *Service
}

func NewHandler(log *gofx.Logger, svc *Service) *Handler {
	return &Handler{log, svc}
}

func (h *Handler) Index(e echo.Context) error {
	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.RenderView(http.StatusOK, page.Device())
}

func (h *Handler) CreateDeviceForm(e echo.Context) error {
	var err error

	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if !c.IsHtmx() {
		return echo.NewHTTPError(http.StatusMethodNotAllowed)
	}

	p := new(CreateDeviceDto)
	if err = c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if err = c.Validate(p); err != nil {
		return c.RenderView(
			http.StatusUnprocessableEntity,
			component.Toast(component.ToastProps{
				Message: err.Error(),
				Type:    "failed",
			}),
		)
	}

	if err = h.svc.CreateDevice(c, p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.RenderView(
		http.StatusOK,
		component.Toast(component.ToastProps{
			Message: "Device created successfully!",
			Type:    "success",
		}),
	)
}

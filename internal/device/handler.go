package device

import (
	"errors"
	"net/http"

	"github.com/hrz8/goatsapp/views/component"
	"github.com/hrz8/goatsapp/views/page"
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

	ck, err := c.Cookie("project_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	projectID := ck.Value
	devices, _ := h.svc.GetDevicesByProjectID(c, projectID)

	return c.RenderView(http.StatusOK, page.Device(devices))
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

func (h *Handler) RequestQR(e echo.Context) error {
	var err error
	var qr string

	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if !c.IsHtmx() {
		return echo.NewHTTPError(http.StatusMethodNotAllowed)
	}

	p := new(RequestQRDto)
	if err = c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	qr, err = h.svc.RequestQRCode(c, p.SessionID)
	if err != nil && errors.Is(err, ErrAlreadyLoggedIn) {
		return c.RenderView(http.StatusOK, page.DeviceQRCode(true, ""))
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.RenderView(http.StatusOK, page.DeviceQRCode(false, qr))
}

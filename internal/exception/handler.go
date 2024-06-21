package exception

import (
	"net/http"

	"github.com/hrz8/goatsapp/web/template/exception"
	"github.com/hrz8/gofx"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) NotFound(c echo.Context) error {
	cc, ok := c.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return cc.RenderView(http.StatusOK, exception.NotFound())
}

package homepage

import (
	"net/http"

	"github.com/hrz8/goatsapp/views/page"
	"github.com/hrz8/gofx"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Index(e echo.Context) error {
	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.RenderView(http.StatusOK, page.Home())
}

package homepage

import (
	"net/http"

	"github.com/hrz8/goatsapp/web/template/page"
	"github.com/hrz8/gofx"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Index(c echo.Context) error {
	cc, ok := c.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return cc.RenderView(http.StatusOK, page.Home())
}

package homepage

import (
	"github.com/hrz8/goatsapp/web/template/exception"
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
		return cc.RenderView(exception.InternalServerError())
	}
	return cc.RenderView(page.Home())
}

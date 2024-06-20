package static

import (
	"embed"

	"github.com/hrz8/goatsapp/web/template/exception"
	"github.com/hrz8/gofx"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Serve(fs embed.FS) func(c echo.Context) error {
	return func(c echo.Context) error {
		cc, ok := c.(*gofx.Context)
		if !ok {
			return cc.RenderView(exception.InternalServerError())
		}
		filePath := c.Request().URL.Path[len("/assets/"):]
		file, err := fs.Open("public/" + filePath)
		if err != nil {
			return cc.RenderView(exception.NotFound())
		}
		defer file.Close()

		return cc.ServeStatic(filePath, file)
	}
}

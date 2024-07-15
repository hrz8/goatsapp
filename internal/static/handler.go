package static

import (
	"embed"
	"net/http"

	"github.com/hrz8/goatsapp/views/exception"
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
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		filePath := c.Request().URL.Path[len("/assets/"):]
		file, err := fs.Open("public/" + filePath)
		if err != nil {
			return cc.RenderView(http.StatusOK, exception.NotFound())
		}
		defer file.Close()

		return cc.ServeStatic(filePath, file)
	}
}

package middleware

import (
	"net/http"

	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Module("middleware", fx.Provide(NewHandler))

type Handler struct {
	repo *dbrepo.Queries
}

func NewHandler(repo *dbrepo.Queries) *Handler {
	return &Handler{repo}
}

func (m *Handler) AppIDCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		_, err = c.Cookie("app_id")
		if err == nil {
			// cookie already set
			return next(c)
		}

		// set cookies
		proj, err := m.repo.GetDefaultProject(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		cookie := new(http.Cookie)
		cookie.Name = "app_id"
		cookie.Value = proj.EncodedID
		cookie.Path = "/"
		c.SetCookie(cookie)

		return next(c)
	}
}

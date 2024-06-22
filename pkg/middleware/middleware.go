package middleware

import (
	"net/http"

	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/labstack/echo/v4"
)

func AppIDCookie(repo *dbrepo.Queries) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			_, err = c.Cookie("app_id")
			if err == nil {
				// cookie already set
				return next(c)
			}

			// set cookies
			proj, err := repo.GetDefaultProject(c.Request().Context())
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			cookie := new(http.Cookie)
			cookie.Name = "app_id"
			cookie.Value = proj.EncodedID
			c.SetCookie(cookie)

			return next(c)
		}
	}
}

package server

import (
	"net/http"

	Assets "github.com/hrz8/goatsapp/assets"
	"github.com/labstack/echo/v4"
)

func index() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(200, "OK")
	}
}

func assets() func(c echo.Context) error {
	return echo.WrapHandler(http.HandlerFunc(Assets.PublicFile))
}

func notFound() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(200, "OK")
	}
}

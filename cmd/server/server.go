package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/hrz8/goatsapp/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Logger *logger.Logger
	HTTP   *echo.Echo
}

func (s *Server) Setup() {
	s.HTTP.GET("/", index())
	s.HTTP.GET("/assets/*", assets())

	s.HTTP.RouteNotFound("", notFound())
	s.HTTP.RouteNotFound("/*", notFound())
}

func (s *Server) Start() <-chan error {
	errCh := make(chan error)

	go func() {
		if err := s.HTTP.Start(":1432"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	return errCh
}

func (s *Server) Cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.HTTP.Shutdown(ctx); err != nil {
		s.Logger.Log.Error("cannot shutdown server", slog.String("err", err.Error()))
		return
	}

	s.Logger.Log.Info("gracefully shutdown http server")
}

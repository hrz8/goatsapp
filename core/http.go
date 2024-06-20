package core

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, logger *Logger, router *Router, cfg *Config) *http.Server {
	logger.Info("registering http server", slog.String("addr", cfg.Addr))
	srv := &http.Server{
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Addr:              cfg.Addr,
		Handler:           router.Mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				err := srv.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("error starting http server", "err", err)
					os.Exit(1)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			toCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(toCtx); err != nil {
				logger.Error("can't shutdown http server gracefully", slog.String("err", err.Error()))
				return err
			}
			logger.Info("gracefully shutdown http server")
			return nil
		},
	})

	return srv
}

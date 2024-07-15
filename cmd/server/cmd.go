package server

import (
	"net/http"

	"github.com/hrz8/goatsapp/config"
	"github.com/hrz8/goatsapp/internal/device"
	"github.com/hrz8/goatsapp/internal/exception"
	"github.com/hrz8/goatsapp/internal/homepage"
	"github.com/hrz8/goatsapp/internal/middleware"
	"github.com/hrz8/goatsapp/internal/project"
	"github.com/hrz8/goatsapp/internal/static"
	"github.com/hrz8/goatsapp/internal/wa"
	"github.com/hrz8/gofx"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "serve [no options!]",
	Short: "Start application server",
	RunE:  run,
}

func run(_ *cobra.Command, _ []string) error {
	cfg := config.New()
	app := gofx.NewApp(&gofx.Config{
		Version:  cfg.AppVersion,
		Addr:     "localhost:" + cfg.AppPort,
		LogLevel: cfg.LogLevel,
	})

	app.AddProviders(NewDB, NewDBRepo, NewWaCli)
	app.AddModules(
		middleware.Module,
		homepage.Module,
		device.Module,
		static.Module,
		exception.Module,
		project.Module,
		wa.Module,
	)
	app.AddProviders(config.New, RegisterRouters, wa.RegisterHandlers)
	app.AddServers(gofx.NewHTTPServer)
	app.AddInvokers(func(*http.Server) {})

	app.Start()

	return nil
}

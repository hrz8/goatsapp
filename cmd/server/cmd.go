package server

import (
	"net/http"

	"github.com/hrz8/goatsapp/config"
	"github.com/hrz8/goatsapp/internal/exception"
	"github.com/hrz8/goatsapp/internal/homepage"
	"github.com/hrz8/goatsapp/internal/project"
	"github.com/hrz8/goatsapp/internal/static"
	"github.com/hrz8/gofx"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "serve [no options!]",
	Short: "Start application server",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	cfg := config.New()

	app := gofx.NewApp(&gofx.Config{
		Version:  cfg.AppVersion,
		Addr:     "localhost:" + cfg.AppPort,
		LogLevel: cfg.LogLevel,
	})

	app.AddProviders(NewDB, NewDBRepo)
	app.AddModules(
		homepage.Module,
		static.Module,
		exception.Module,
		project.Module,
	)
	app.AddProviders(RegisterRouters, config.New)
	app.AddServers(gofx.NewHTTPServer)
	app.AddInvokers(func(*http.Server) {})

	app.Start()
}

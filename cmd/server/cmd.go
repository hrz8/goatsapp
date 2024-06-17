package server

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/hrz8/goatsapp/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server [no options!]",
	Short: "Start application server",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	lgr := logger.NewLogger()
	e := echo.New()

	server := &Server{Logger: lgr, HTTP: e}
	server.Setup()

	serverErr := server.Start()
	defer server.Cleanup()

	select {
	case sig := <-waitShutdown():
		lgr.Log.Error(fmt.Sprintf("%s signal triggered", sig))
	case err := <-serverErr:
		lgr.Log.Error("cannot start the server", slog.String("err", err.Error()))
	}
}

func waitShutdown() <-chan os.Signal {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	return exit
}

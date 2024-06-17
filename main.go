package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/hrz8/goatsapp/cmd/server"
	"github.com/hrz8/goatsapp/cmd/version"
	"github.com/hrz8/goatsapp/config"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "goatsapp",
	Version: config.Version,
	Short:   "WhatsApp Unofficial API",
	Long:    `An open source application of unofficial WhatsApp API based on websocket`,
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println("Welcome to " + cmd.Use + "@" + cmd.Version) //nolint
	},
}

func main() {
	cmd.AddCommand(version.VersionCmd)
	cmd.AddCommand(server.ServerCmd)

	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

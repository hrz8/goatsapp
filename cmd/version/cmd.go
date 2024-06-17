package version

import (
	"fmt"

	"github.com/hrz8/goatsapp/config"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version [no options!]",
	Short: "Print the version number of the application",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	fmt.Println(config.Version) //nolint
}

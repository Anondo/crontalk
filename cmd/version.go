package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "displays the app version",
		Run:   showVersion,
	}
)

func showVersion(cmd *cobra.Command, args []string) {
	v := viper.GetString("app.version")
	fmt.Printf("CronTalk version %s\n", v)
}

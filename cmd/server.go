package cmd

import (
	"github.com/Anondo/crontalk/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the crontalk http server",
		Run:   serve,
	}
)

func init() {
	serveCmd.Flags().IntP("port", "p", viper.GetInt("app.port"),
		"The http port to start the crontalk server on")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}

func serve(cmd *cobra.Command, args []string) {
	server.StartServer()
}

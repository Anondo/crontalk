package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// Root Command
var (
	RootCmd = &cobra.Command{
		Use:   "crontalk",
		Short: "crontalk ...",
	}
)

func init() {
	RootCmd.AddCommand(translateCmd)
}

// Execute the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

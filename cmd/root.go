package cmd

import (
	"log"

	"github.com/Anondo/crontalk/config"
	"github.com/Anondo/crontalk/helper"

	"github.com/spf13/cobra"
)

// Root Command
var (
	RootCmd = &cobra.Command{
		Use:   "crontalk",
		Short: "Talk to the cron expressions!!",
		Long: helper.Logo + `
CronTalk: Talk to the cron expressions!!
------------------------------------------

Translate your cron expressions into very understandable english words.
Scan the expression for future occurrences.

Valid Cron Expressions:
-----------------------
  "* * * * *" = "(minute) (hour) (day of month) (month) (day of week)"
1.Should contain exactly 5 values/sub-expressions
2. Valid values are:
		minute: 0-59
		hour:   0-23
		day of month: 1-31
		month: 1-12 or jan-dec
		day of week: 0-6 or sun-sat
		`,
	}
)

func init() {
	config.LoadConfig()
	RootCmd.AddCommand(translateCmd)
	RootCmd.AddCommand(nextCmd)
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(generateCmd)
}

// Execute the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

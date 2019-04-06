package cmd

import (
	"crontalk/helper"
	"log"

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
Scan the expression for previous/future occurences.

Valid Cron Expressions:
-----------------------
  "* * * * *" = "(minute) (hour) (day of month) (month) (day of week)"
1.Should contain exactly 5 values/sub-expressions
2.Values ar only limited to numeric digits for now

		`,
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

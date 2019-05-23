package cmd

import (
	"fmt"
	"time"

	translator "crontalk/translator.go"
	"github.com/gorhill/cronexpr"
	"github.com/spf13/cobra"
)

var (
	nextCmd = &cobra.Command{
		Use:   "next",
		Short: "Shows the next occurrence of a cron expression",
		Run:   occur,
	}
	occurenceNumber = 1
	layout          = "2006-01-02 03:04PM"
)

func init() {
	nextCmd.Flags().StringVarP(&translator.CronExprsn, "cron", "c", "", "The cron expression to scan for occurrence")
	nextCmd.Flags().IntVarP(&occurenceNumber, "occurence", "o", 1, "The number of occurence time")

}

func occur(cmd *cobra.Command, args []string) {

	if vErr := translator.Validate(); len(vErr) != 0 {
		for k, v := range vErr {
			fmt.Printf("%v: %v\n", k, v)
		}
		return
	}
	exprns := cronexpr.MustParse(translator.CronExprsn).NextN(time.Now(), uint(occurenceNumber))
	for _, expr := range exprns {
		fmt.Println(expr.Format(layout))
	}
}

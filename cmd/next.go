package cmd

import (
	"fmt"
	"log"
	"time"

	translator "crontalk/translator"

	"github.com/gorhill/cronexpr"
	"github.com/spf13/cobra"
)

var (
	nextCmd = &cobra.Command{
		Use:     "next",
		Short:   "Shows the next occurrence of a cron expression",
		Example: `crontalk next "* * * * *"`,
		Run:     occur,
	}
	occurenceNumber = 1
	layout          = "2006-01-02 03:04PM"
)

func init() {
	nextCmd.Flags().IntVarP(&occurenceNumber, "occurence", "o", 1, "The number of occurence time")

}

func occur(cmd *cobra.Command, args []string) {

	if len(args) < 1 {
		log.Fatal("no cron expression detected")
	}

	translator.CronExprsn = args[0]

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

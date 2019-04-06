package cmd

import (
	"fmt"
	"log"

	"crontalk/translator.go"
	"github.com/spf13/cobra"
)

var (
	translateCmd = &cobra.Command{
		Use:   "translate",
		Short: "translates the given cron expression to english words",
		Run:   translate,
	}
)

func init() {
	translateCmd.Flags().StringVarP(&translator.CronExprsn, "cron", "c", "", "the cron expression to translate to english words")

}

func translate(cmd *cobra.Command, args []string) {

	if vErr := translator.Validate(); len(vErr) != 0 {
		for k, v := range vErr {
			fmt.Printf("%v: %v\n", k, v)
		}
		return
	}

	if err := translator.Translate(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(translator.GetTranslatedStr())

}

package cmd

import (
	"fmt"
	"log"

	"crontalk/helper"

	"crontalk/translator"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	translateCmd = &cobra.Command{
		Use:   "translate",
		Short: "Translates the given cron expression to english words",
		Run:   translate,
	}
)

func init() {
	translateCmd.Flags().StringVarP(&translator.CronExprsn, "cron", "c", "", "The cron expression to translate to english words")
	translateCmd.Flags().BoolP("bangla", "b", false, "The translation to be in Bangla language")
	viper.BindPFlag("bangla", translateCmd.Flags().Lookup("bangla"))
}

func translate(cmd *cobra.Command, args []string) {

	translator.Init()

	if vErr := translator.Validate(); len(vErr) != 0 {
		for k, v := range vErr {
			fmt.Printf("%v: %v\n", k, v)
		}
		return
	}

	if err := translator.Translate(); err != nil {
		log.Fatal(err.Error())
	}

	output := translator.GetTranslatedStr()
	output = helper.TrimExtraSpaces(output)

	if viper.GetBool("bangla") {
		helper.ChangeDigitLanguage(&output, "bangla") //changing the english digits to bangla
	}

	output = helper.AddOrdinals(output)

	fmt.Println(output)

}

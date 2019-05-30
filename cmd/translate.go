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
		Use:     "translate",
		Short:   "Translates the given cron expression to english words",
		Example: `crontalk translate "* * * * *"`,
		Run:     translate,
	}
)

func init() {
	translateCmd.Flags().BoolP("bangla", "b", false, "The translation to be in Bangla language")
	viper.BindPFlag("bangla", translateCmd.Flags().Lookup("bangla"))
}

func translate(cmd *cobra.Command, args []string) {

	if len(args) < 1 {
		log.Fatal("no cron expression detected")
	}

	translator.CronExprsn = args[0]

	translator.Init()

	if vErr := translator.Validate(); len(vErr) != 0 {
		for en, ev := range vErr {
			fmt.Printf("%v:\n", en)
			for i, e := range ev {
				fmt.Printf("%d.%v\n", i+1, e)
			}
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

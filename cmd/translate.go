package cmd

import (
	"fmt"
	"strings"

	"github.com/Anondo/crontalk/helper"

	"github.com/Anondo/crontalk/translator"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	translateCmd = &cobra.Command{
		Use:     "translate",
		Short:   "Translates the given cron expression to english words",
		RunE:    translate,
		Example: `crontalk translate "* * * * *"`,
		Args:    cobra.ExactArgs(1),
	}
)

func init() {
	translateCmd.Flags().StringP("lang", "l", "english", "Set the translation language")
	viper.BindPFlag("lang", translateCmd.Flags().Lookup("lang"))
}

func translate(cmd *cobra.Command, args []string) error {

	lang := strings.ToLower(viper.GetString("lang"))

	if lang != helper.LanguageEnglish {
		langMap := viper.GetStringMap("language")
		if _, exists := langMap[lang]; !exists {
			possibleLangs := []string{}
			for pl := range langMap {
				possibleLangs = append(possibleLangs, pl)
			}
			return fmt.Errorf("%s, valid values: %v", helper.ErrInvalidLangValue, possibleLangs)
		}
	}

	translator.Init()

	tr := translator.NewTranslator(args[0])

	if vErr := tr.Validate(); len(vErr) != 0 {
		for en, ev := range vErr {
			fmt.Printf("%v:\n", en)
			for i, e := range ev {
				fmt.Printf("%d.%v\n", i+1, e)
			}
		}
		return nil
	}

	if err := tr.Translate(); err != nil {
		return err
	}

	output := tr.GetTranslatedStr()
	output = helper.TrimExtraSpaces(output)

	if lang != helper.LanguageEnglish {
		helper.ChangeDigitLanguage(&output, lang) //changing the english digits to different language
	}

	output = helper.AddOrdinals(output)

	fmt.Println(output)

	return nil

}

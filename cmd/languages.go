package cmd

import (
	"fmt"

	"github.com/Anondo/crontalk/config"
	"github.com/Anondo/crontalk/helper"
	"github.com/spf13/cobra"
)

var (
	languagesCmd = &cobra.Command{
		Use:     "languages",
		Short:   "Shows a list of supported languages for translation",
		Example: `crontalk languages`,
		Run:     showLangs,
	}
)

func showLangs(cmd *cobra.Command, args []string) {

	langMap := config.LanguageMap()
	no := 1
	for lang := range langMap {
		fmt.Printf("%d. %s", no, lang)
		if lang == helper.LanguageEnglish {
			fmt.Printf("(default)")
		}
		fmt.Println()

		no++
	}

}

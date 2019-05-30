package cmd

import (
	"crontalk/generator"
	translator "crontalk/translator"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "generates a cron expression from english words by prompting the user",
		RunE:  generate,
	}
)

func generate(cmd *cobra.Command, args []string) error {
	expr, err := generator.GenerateCron()
	if err != nil {
		return err
	}
	fmt.Println("The cron expression:", expr)
	fmt.Print("Translation: ")
	translator.CronExprsn = expr
	translate(cmd, []string{})

	return err
}

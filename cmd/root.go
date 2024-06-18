package cmd

import (
	"os"

	"github.com/GoToolSharing/passfinder/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passfinder",
	Short: "Generate password wordlists with customizable options for security testing !",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&config.GlobalConfig.OutputFile, "output", "o", "", "Write to output file")
	rootCmd.PersistentFlags().BoolVar(&config.GlobalConfig.BatchParam, "batch", false, "Don't ask questions")
}

package cmd

import (
	"os"

	"github.com/GoToolSharing/passfinder/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passfinder",
	Short: "TODO",
	Long:  `TODO`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&config.GlobalConfig.OutputFile, "output", "", "", "Output file")
}

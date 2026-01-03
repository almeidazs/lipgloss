package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: false,
	Use:           "gloss <command> [flags]",
	Short:         "A stateful and stylish HTTP CLI",
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

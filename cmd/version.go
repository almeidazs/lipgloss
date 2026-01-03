package cmd

import (
	"fmt"

	"github.com.almeidazs/lipgloss/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Shows the current version of Lipgloss",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Println(version.Version)

		return err
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

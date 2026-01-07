package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com.almeidazs/lipgloss/internal/color"
	"github.com/creativeprojects/go-selfupdate"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Lipgloss to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		return update(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func update(ctx context.Context) error {
	if rootCmd.Version == "" {
		return errors.New("current version is not available")
	}

	ctx, Cancel := context.WithTimeout(ctx, 10*time.Second)

	defer Cancel()

	const slug = "almeidazs/lipgloss"

	latest, found, err := selfupdate.DetectLatest(ctx, selfupdate.ParseSlug(slug))

	if err != nil {
		return fmt.Errorf("failed to check latest version: %w", err)
	}

	if !found || latest.LessOrEqual(rootCmd.Version) {
		fmt.Printf("You're already running the latest version (%s)\n", color.Paint(color.Gray, rootCmd.Version))

		return nil
	}

	exe, err := os.Executable()

	if err != nil {
		return fmt.Errorf("failed to locate executable: %w", err)
	}

	fmt.Printf("Updating Lipgloss from %s to %s...\n", color.Paint(color.Gray, rootCmd.Version), color.Paint(color.Gray, latest.Version()))

	if err := selfupdate.UpdateTo(ctx, latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	fmt.Printf("You'are now on the %s of Lipgloss!\n", color.Paint(color.Gray, latest.Version()))

	return nil
}

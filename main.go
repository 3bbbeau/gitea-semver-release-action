package main

import (
	"os"

	"github.com/3bbbeau/gitea-semver-release-action/internal/pkg/event"
	"github.com/3bbbeau/gitea-semver-release-action/internal/pkg/git"
	"github.com/3bbbeau/gitea-semver-release-action/internal/pkg/release"
	"github.com/3bbbeau/gitea-semver-release-action/internal/pkg/semver"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "bumper",
	}
	rootCmd.SetOut(os.Stdout)

	rootCmd.AddCommand(semver.Command())
	rootCmd.AddCommand(release.Command())
	rootCmd.AddCommand(event.GuardCommand())
	rootCmd.AddCommand(event.IncrementCommand())
	rootCmd.AddCommand(git.LatestTagCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

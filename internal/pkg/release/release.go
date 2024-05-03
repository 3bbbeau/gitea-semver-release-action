package release

import (
	"strings"

	"code.gitea.io/sdk/gitea"
	"github.com/3bbbeau/gitea-semver-release-action/internal/pkg/action"
	"github.com/spf13/cobra"
)

const releaseTypeNone = "none"
const releaseTypeRelease = "release"
const releaseTypeTag = "tag"

type repository struct {
	owner string
	name  string
	token string
}

type releaseDetails struct {
	version string
	target  string
}

func Command() *cobra.Command {
	var releaseType string

	cmd := &cobra.Command{
		Use:  "release [REPOSITORY] [TARGET_COMMITISH] [VERSION] [GH_TOKEN]",
		Args: cobra.ExactArgs(5),
		Run: func(cmd *cobra.Command, args []string) {
			execute(cmd, releaseType, args)
		},
	}

	cmd.Flags().StringVarP(&releaseType, "strategy", "s", releaseTypeRelease, "Release strategy")

	return cmd
}

func execute(cmd *cobra.Command, releaseType string, args []string) {
	parts := strings.Split(args[0], "/")
	repo := repository{
		owner: parts[0],
		name:  parts[1],
		token: args[3],
	}

	release := releaseDetails{
		version: args[2],
		target:  args[1],
	}

	apiURL := args[4]
	client, err := gitea.NewClient(apiURL, gitea.SetToken(repo.token))
	if err != nil {
		action.Fail(cmd, "could not create gitea client: %s", err)
	}

	switch releaseType {
	case releaseTypeNone:
		return
	case releaseTypeRelease:
		if err := createGithubRelease(client, repo, release); err != nil {
			action.AssertNoError(cmd, err, "could not create GitHub release: %s", err)
		}
		return
	case releaseTypeTag:
		if err := createLightweightTag(client, repo, release); err != nil {
			action.AssertNoError(cmd, err, "could not create lightweight tag: %s", err)
		}
		return
	default:
		action.Fail(cmd, "unknown release strategy: %s", releaseType)
	}
}

func createLightweightTag(client *gitea.Client, repo repository, release releaseDetails) error {
	_, _, err := client.CreateTag(repo.owner, repo.name, gitea.CreateTagOption{
		TagName: release.version,
		Target:  release.target,
	})

	return err
}

func createGithubRelease(client *gitea.Client, repo repository, release releaseDetails) error {
	_, _, err := client.CreateRelease(repo.owner, repo.name, gitea.CreateReleaseOption{
		Title:        release.version,
		TagName:      release.version,
		Target:       release.target,
		IsDraft:      false,
		IsPrerelease: false,
	})

	return err
}

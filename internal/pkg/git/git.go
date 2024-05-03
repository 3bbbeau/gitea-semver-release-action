package git

import (
	"net/http"
	"strings"

	"code.gitea.io/sdk/gitea"
	"github.com/3bbbeau/gitea-semver-release-action/internal/pkg/action"
	"github.com/blang/semver/v4"
	"github.com/spf13/cobra"
)

func LatestTagCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "latest-tag [REPOSITORY] [GH_TOKEN]",
		Args: cobra.ExactArgs(3),
		Run:  executeLatestTag,
	}
}

func executeLatestTag(cmd *cobra.Command, args []string) {
	repository := args[0]
	token := args[1]
	apiURL := args[2]

	client, err := gitea.NewClient(apiURL, gitea.SetToken(token))
	if err != nil {
		action.Fail(cmd, "could not create gitea client: %s", err)
	}

	parts := strings.Split(repository, "/")
	owner := parts[0]
	repo := parts[1]

	refs, response, err := client.GetRepoRefs(owner, repo, "tags")
	if response != nil && response.StatusCode == http.StatusNotFound {
		cmd.Print("v0.0.0")
		return
	}
	action.AssertNoError(cmd, err, "could not list git refs: %s", err)

	latest := semver.MustParse("0.0.0")
	for _, ref := range refs {
		version, err := semver.ParseTolerant(strings.Replace(ref.Ref, "refs/tags/", "", 1))
		if err != nil {
			continue
		}

		if version.GT(latest) {
			latest = version
		}
	}

	cmd.Printf("v%s", latest)
}

# (Fork) Semver Release Gitea Action

> [!NOTE]  
> Changes in this fork allow this action to work on Gitea. I was too lazy to abstract the codebase to accommodate both clients, so this changeset can be taken to bring that functionality upstream by someone who would like to do so.

Automatically create [SemVer](https://semver.org/) compliant releases based on
PR labels.

Assuming that a PR is tagged with a "*semver-compliant*" label (*patch*, *minor* or *major*),
then this action can create a tag and a GitHub release when it is merged.

**Note:** to determine the base tag for the increment, this action will try to
find the most recent tag complying to [SemVer](https://semver.org/). No
additional setup is required.

## Inputs

### `release_branch`

**Required** Branch to tag. Default `"master"`.

### `release_strategy`

**Required** Release strategy. Default `"release"` (`release`: creates a GitHub
release ; `tag`: creates a lightweight tag ; `none`: computes the next
[SemVer](https://semver.org/) version but does not create a release or tag).

### `tag_format`

**Optional** Format used to create tags. Default `"v%major%.%minor%.%patch%"`.

### `tag`

**Optional** Tag to use. If left undefined, it will be computed using the tags
already present in the repository.

### `api_url`
**Optional**  Your Gitea instance, without the `/v1/api`. Defaults to `https://gitea.com`

## Outputs

### `tag`

The newly created tag.

## Example usage

```yaml
# .github/workflows/release.yml
name: Release

on:
  pull_request:
    types: [closed]

jobs:
  build:
    runs-on: ubuntu-latest

    if: github.event.pull_request.merged
    
    steps:
      - name: Tag
        uses: 3bbbeau/gitea-semver-release-action@master
        with:
          release_branch: master
          api_url: "https://gitea.mydomain.tld" 
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

```

## License

This library is under the [MIT](LICENSE.md) license.

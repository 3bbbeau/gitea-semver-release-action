package event

import "time"

// Gitea PullRequest event
//
// SDK doesn't include all fields, hence it is defined hre.
type PullRequestEvent struct {
	Action      string      `json:"action,omitempty"`
	Number      int         `json:"number,omitempty"`
	PullRequest PullRequest `json:"pull_request,omitempty"`
	Repository  Repository  `json:"repository,omitempty"`
	Sender      Sender      `json:"sender,omitempty"`
	CommitID    string      `json:"commit_id,omitempty"`
	Review      any         `json:"review,omitempty"`
}
type User struct {
	ID                int       `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	LoginName         string    `json:"login_name,omitempty"`
	FullName          string    `json:"full_name,omitempty"`
	Email             string    `json:"email,omitempty"`
	AvatarURL         string    `json:"avatar_url,omitempty"`
	Language          string    `json:"language,omitempty"`
	IsAdmin           bool      `json:"is_admin,omitempty"`
	LastLogin         time.Time `json:"last_login,omitempty"`
	Created           time.Time `json:"created,omitempty"`
	Restricted        bool      `json:"restricted,omitempty"`
	Active            bool      `json:"active,omitempty"`
	ProhibitLogin     bool      `json:"prohibit_login,omitempty"`
	Location          string    `json:"location,omitempty"`
	Website           string    `json:"website,omitempty"`
	Description       string    `json:"description,omitempty"`
	Visibility        string    `json:"visibility,omitempty"`
	FollowersCount    int       `json:"followers_count,omitempty"`
	FollowingCount    int       `json:"following_count,omitempty"`
	StarredReposCount int       `json:"starred_repos_count,omitempty"`
	Username          string    `json:"username,omitempty"`
}
type Labels struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}
type Milestone struct {
	ID           int       `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	State        string    `json:"state,omitempty"`
	OpenIssues   int       `json:"open_issues,omitempty"`
	ClosedIssues int       `json:"closed_issues,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	ClosedAt     any       `json:"closed_at,omitempty"`
	DueOn        any       `json:"due_on,omitempty"`
}
type Assignee struct {
	ID                int       `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	LoginName         string    `json:"login_name,omitempty"`
	FullName          string    `json:"full_name,omitempty"`
	Email             string    `json:"email,omitempty"`
	AvatarURL         string    `json:"avatar_url,omitempty"`
	Language          string    `json:"language,omitempty"`
	IsAdmin           bool      `json:"is_admin,omitempty"`
	LastLogin         time.Time `json:"last_login,omitempty"`
	Created           time.Time `json:"created,omitempty"`
	Restricted        bool      `json:"restricted,omitempty"`
	Active            bool      `json:"active,omitempty"`
	ProhibitLogin     bool      `json:"prohibit_login,omitempty"`
	Location          string    `json:"location,omitempty"`
	Website           string    `json:"website,omitempty"`
	Description       string    `json:"description,omitempty"`
	Visibility        string    `json:"visibility,omitempty"`
	FollowersCount    int       `json:"followers_count,omitempty"`
	FollowingCount    int       `json:"following_count,omitempty"`
	StarredReposCount int       `json:"starred_repos_count,omitempty"`
	Username          string    `json:"username,omitempty"`
}
type Assignees struct {
	ID                int       `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	LoginName         string    `json:"login_name,omitempty"`
	FullName          string    `json:"full_name,omitempty"`
	Email             string    `json:"email,omitempty"`
	AvatarURL         string    `json:"avatar_url,omitempty"`
	Language          string    `json:"language,omitempty"`
	IsAdmin           bool      `json:"is_admin,omitempty"`
	LastLogin         time.Time `json:"last_login,omitempty"`
	Created           time.Time `json:"created,omitempty"`
	Restricted        bool      `json:"restricted,omitempty"`
	Active            bool      `json:"active,omitempty"`
	ProhibitLogin     bool      `json:"prohibit_login,omitempty"`
	Location          string    `json:"location,omitempty"`
	Website           string    `json:"website,omitempty"`
	Description       string    `json:"description,omitempty"`
	Visibility        string    `json:"visibility,omitempty"`
	FollowersCount    int       `json:"followers_count,omitempty"`
	FollowingCount    int       `json:"following_count,omitempty"`
	StarredReposCount int       `json:"starred_repos_count,omitempty"`
	Username          string    `json:"username,omitempty"`
}
type MergedBy struct {
	ID                int       `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	LoginName         string    `json:"login_name,omitempty"`
	FullName          string    `json:"full_name,omitempty"`
	Email             string    `json:"email,omitempty"`
	AvatarURL         string    `json:"avatar_url,omitempty"`
	Language          string    `json:"language,omitempty"`
	IsAdmin           bool      `json:"is_admin,omitempty"`
	LastLogin         time.Time `json:"last_login,omitempty"`
	Created           time.Time `json:"created,omitempty"`
	Restricted        bool      `json:"restricted,omitempty"`
	Active            bool      `json:"active,omitempty"`
	ProhibitLogin     bool      `json:"prohibit_login,omitempty"`
	Location          string    `json:"location,omitempty"`
	Website           string    `json:"website,omitempty"`
	Description       string    `json:"description,omitempty"`
	Visibility        string    `json:"visibility,omitempty"`
	FollowersCount    int       `json:"followers_count,omitempty"`
	FollowingCount    int       `json:"following_count,omitempty"`
	StarredReposCount int       `json:"starred_repos_count,omitempty"`
	Username          string    `json:"username,omitempty"`
}
type Owner struct {
	ID                int       `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	LoginName         string    `json:"login_name,omitempty"`
	FullName          string    `json:"full_name,omitempty"`
	Email             string    `json:"email,omitempty"`
	AvatarURL         string    `json:"avatar_url,omitempty"`
	Language          string    `json:"language,omitempty"`
	IsAdmin           bool      `json:"is_admin,omitempty"`
	LastLogin         time.Time `json:"last_login,omitempty"`
	Created           time.Time `json:"created,omitempty"`
	Restricted        bool      `json:"restricted,omitempty"`
	Active            bool      `json:"active,omitempty"`
	ProhibitLogin     bool      `json:"prohibit_login,omitempty"`
	Location          string    `json:"location,omitempty"`
	Website           string    `json:"website,omitempty"`
	Description       string    `json:"description,omitempty"`
	Visibility        string    `json:"visibility,omitempty"`
	FollowersCount    int       `json:"followers_count,omitempty"`
	FollowingCount    int       `json:"following_count,omitempty"`
	StarredReposCount int       `json:"starred_repos_count,omitempty"`
	Username          string    `json:"username,omitempty"`
}
type Permissions struct {
	Admin bool `json:"admin,omitempty"`
	Push  bool `json:"push,omitempty"`
	Pull  bool `json:"pull,omitempty"`
}
type InternalTracker struct {
	EnableTimeTracker                bool `json:"enable_time_tracker,omitempty"`
	AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time,omitempty"`
	EnableIssueDependencies          bool `json:"enable_issue_dependencies,omitempty"`
}
type Repo struct {
	ID                            int             `json:"id,omitempty"`
	Owner                         Owner           `json:"owner,omitempty"`
	Name                          string          `json:"name,omitempty"`
	FullName                      string          `json:"full_name,omitempty"`
	Description                   string          `json:"description,omitempty"`
	Empty                         bool            `json:"empty,omitempty"`
	Private                       bool            `json:"private,omitempty"`
	Fork                          bool            `json:"fork,omitempty"`
	Template                      bool            `json:"template,omitempty"`
	Parent                        any             `json:"parent,omitempty"`
	Mirror                        bool            `json:"mirror,omitempty"`
	Size                          int             `json:"size,omitempty"`
	Language                      string          `json:"language,omitempty"`
	LanguagesURL                  string          `json:"languages_url,omitempty"`
	HTMLURL                       string          `json:"html_url,omitempty"`
	SSHURL                        string          `json:"ssh_url,omitempty"`
	CloneURL                      string          `json:"clone_url,omitempty"`
	OriginalURL                   string          `json:"original_url,omitempty"`
	Website                       string          `json:"website,omitempty"`
	StarsCount                    int             `json:"stars_count,omitempty"`
	ForksCount                    int             `json:"forks_count,omitempty"`
	WatchersCount                 int             `json:"watchers_count,omitempty"`
	OpenIssuesCount               int             `json:"open_issues_count,omitempty"`
	OpenPrCounter                 int             `json:"open_pr_counter,omitempty"`
	ReleaseCounter                int             `json:"release_counter,omitempty"`
	DefaultBranch                 string          `json:"default_branch,omitempty"`
	Archived                      bool            `json:"archived,omitempty"`
	CreatedAt                     time.Time       `json:"created_at,omitempty"`
	UpdatedAt                     time.Time       `json:"updated_at,omitempty"`
	Permissions                   Permissions     `json:"permissions,omitempty"`
	HasIssues                     bool            `json:"has_issues,omitempty"`
	InternalTracker               InternalTracker `json:"internal_tracker,omitempty"`
	HasWiki                       bool            `json:"has_wiki,omitempty"`
	HasPullRequests               bool            `json:"has_pull_requests,omitempty"`
	HasProjects                   bool            `json:"has_projects,omitempty"`
	IgnoreWhitespaceConflicts     bool            `json:"ignore_whitespace_conflicts,omitempty"`
	AllowMergeCommits             bool            `json:"allow_merge_commits,omitempty"`
	AllowRebase                   bool            `json:"allow_rebase,omitempty"`
	AllowRebaseExplicit           bool            `json:"allow_rebase_explicit,omitempty"`
	AllowSquashMerge              bool            `json:"allow_squash_merge,omitempty"`
	AllowRebaseUpdate             bool            `json:"allow_rebase_update,omitempty"`
	DefaultDeleteBranchAfterMerge bool            `json:"default_delete_branch_after_merge,omitempty"`
	DefaultMergeStyle             string          `json:"default_merge_style,omitempty"`
	AvatarURL                     string          `json:"avatar_url,omitempty"`
	Internal                      bool            `json:"internal,omitempty"`
	MirrorInterval                string          `json:"mirror_interval,omitempty"`
	MirrorUpdated                 time.Time       `json:"mirror_updated,omitempty"`
	RepoTransfer                  any             `json:"repo_transfer,omitempty"`
}
type Base struct {
	Label  string `json:"label,omitempty"`
	Ref    string `json:"ref,omitempty"`
	Sha    string `json:"sha,omitempty"`
	RepoID int    `json:"repo_id,omitempty"`
	Repo   Repo   `json:"repo,omitempty"`
}
type Head struct {
	Label  string `json:"label,omitempty"`
	Ref    string `json:"ref,omitempty"`
	Sha    string `json:"sha,omitempty"`
	RepoID int    `json:"repo_id,omitempty"`
	Repo   Repo   `json:"repo,omitempty"`
}
type PullRequest struct {
	ID                  int         `json:"id,omitempty"`
	URL                 string      `json:"url,omitempty"`
	Number              int         `json:"number,omitempty"`
	User                User        `json:"user,omitempty"`
	Title               string      `json:"title,omitempty"`
	Body                string      `json:"body,omitempty"`
	Labels              []Labels    `json:"labels,omitempty"`
	Milestone           Milestone   `json:"milestone,omitempty"`
	Assignee            Assignee    `json:"assignee,omitempty"`
	Assignees           []Assignees `json:"assignees,omitempty"`
	State               string      `json:"state,omitempty"`
	IsLocked            bool        `json:"is_locked,omitempty"`
	Comments            int         `json:"comments,omitempty"`
	HTMLURL             string      `json:"html_url,omitempty"`
	DiffURL             string      `json:"diff_url,omitempty"`
	PatchURL            string      `json:"patch_url,omitempty"`
	Mergeable           bool        `json:"mergeable,omitempty"`
	Merged              bool        `json:"merged,omitempty"`
	MergedAt            time.Time   `json:"merged_at,omitempty"`
	MergeCommitSha      string      `json:"merge_commit_sha,omitempty"`
	MergedBy            MergedBy    `json:"merged_by,omitempty"`
	AllowMaintainerEdit bool        `json:"allow_maintainer_edit,omitempty"`
	Base                Base        `json:"base,omitempty"`
	Head                Head        `json:"head,omitempty"`
	MergeBase           string      `json:"merge_base,omitempty"`
	DueDate             any         `json:"due_date,omitempty"`
	CreatedAt           time.Time   `json:"created_at,omitempty"`
	UpdatedAt           time.Time   `json:"updated_at,omitempty"`
	ClosedAt            time.Time   `json:"closed_at,omitempty"`
}
type Repository struct {
	ID                            int             `json:"id,omitempty"`
	Owner                         Owner           `json:"owner,omitempty"`
	Name                          string          `json:"name,omitempty"`
	FullName                      string          `json:"full_name,omitempty"`
	Description                   string          `json:"description,omitempty"`
	Empty                         bool            `json:"empty,omitempty"`
	Private                       bool            `json:"private,omitempty"`
	Fork                          bool            `json:"fork,omitempty"`
	Template                      bool            `json:"template,omitempty"`
	Parent                        any             `json:"parent,omitempty"`
	Mirror                        bool            `json:"mirror,omitempty"`
	Size                          int             `json:"size,omitempty"`
	Language                      string          `json:"language,omitempty"`
	LanguagesURL                  string          `json:"languages_url,omitempty"`
	HTMLURL                       string          `json:"html_url,omitempty"`
	SSHURL                        string          `json:"ssh_url,omitempty"`
	CloneURL                      string          `json:"clone_url,omitempty"`
	OriginalURL                   string          `json:"original_url,omitempty"`
	Website                       string          `json:"website,omitempty"`
	StarsCount                    int             `json:"stars_count,omitempty"`
	ForksCount                    int             `json:"forks_count,omitempty"`
	WatchersCount                 int             `json:"watchers_count,omitempty"`
	OpenIssuesCount               int             `json:"open_issues_count,omitempty"`
	OpenPrCounter                 int             `json:"open_pr_counter,omitempty"`
	ReleaseCounter                int             `json:"release_counter,omitempty"`
	DefaultBranch                 string          `json:"default_branch,omitempty"`
	Archived                      bool            `json:"archived,omitempty"`
	CreatedAt                     time.Time       `json:"created_at,omitempty"`
	UpdatedAt                     time.Time       `json:"updated_at,omitempty"`
	Permissions                   Permissions     `json:"permissions,omitempty"`
	HasIssues                     bool            `json:"has_issues,omitempty"`
	InternalTracker               InternalTracker `json:"internal_tracker,omitempty"`
	HasWiki                       bool            `json:"has_wiki,omitempty"`
	HasPullRequests               bool            `json:"has_pull_requests,omitempty"`
	HasProjects                   bool            `json:"has_projects,omitempty"`
	IgnoreWhitespaceConflicts     bool            `json:"ignore_whitespace_conflicts,omitempty"`
	AllowMergeCommits             bool            `json:"allow_merge_commits,omitempty"`
	AllowRebase                   bool            `json:"allow_rebase,omitempty"`
	AllowRebaseExplicit           bool            `json:"allow_rebase_explicit,omitempty"`
	AllowSquashMerge              bool            `json:"allow_squash_merge,omitempty"`
	AllowRebaseUpdate             bool            `json:"allow_rebase_update,omitempty"`
	DefaultDeleteBranchAfterMerge bool            `json:"default_delete_branch_after_merge,omitempty"`
	DefaultMergeStyle             string          `json:"default_merge_style,omitempty"`
	AvatarURL                     string          `json:"avatar_url,omitempty"`
	Internal                      bool            `json:"internal,omitempty"`
	MirrorInterval                string          `json:"mirror_interval,omitempty"`
	MirrorUpdated                 time.Time       `json:"mirror_updated,omitempty"`
	RepoTransfer                  any             `json:"repo_transfer,omitempty"`
}
type Sender struct {
	ID                int       `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	LoginName         string    `json:"login_name,omitempty"`
	FullName          string    `json:"full_name,omitempty"`
	Email             string    `json:"email,omitempty"`
	AvatarURL         string    `json:"avatar_url,omitempty"`
	Language          string    `json:"language,omitempty"`
	IsAdmin           bool      `json:"is_admin,omitempty"`
	LastLogin         time.Time `json:"last_login,omitempty"`
	Created           time.Time `json:"created,omitempty"`
	Restricted        bool      `json:"restricted,omitempty"`
	Active            bool      `json:"active,omitempty"`
	ProhibitLogin     bool      `json:"prohibit_login,omitempty"`
	Location          string    `json:"location,omitempty"`
	Website           string    `json:"website,omitempty"`
	Description       string    `json:"description,omitempty"`
	Visibility        string    `json:"visibility,omitempty"`
	FollowersCount    int       `json:"followers_count,omitempty"`
	FollowingCount    int       `json:"following_count,omitempty"`
	StarredReposCount int       `json:"starred_repos_count,omitempty"`
	Username          string    `json:"username,omitempty"`
}

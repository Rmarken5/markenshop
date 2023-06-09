package pull_request

import "time"

type Pull struct {
	Action      string      `json:"action,omitempty"`
	Number      int         `json:"number,omitempty"`
	PullRequest PullRequest `json:"pull_request,omitempty"`
	Repository  Repository  `json:"repository,omitempty"`
	Sender      Sender      `json:"sender,omitempty"`
}
type User struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
type Owner struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
type Repo struct {
	ID                        int       `json:"id,omitempty"`
	NodeID                    string    `json:"node_id,omitempty"`
	Name                      string    `json:"name,omitempty"`
	FullName                  string    `json:"full_name,omitempty"`
	Private                   bool      `json:"private,omitempty"`
	Owner                     Owner     `json:"owner,omitempty"`
	HTMLURL                   string    `json:"html_url,omitempty"`
	Description               string    `json:"description,omitempty"`
	Fork                      bool      `json:"fork,omitempty"`
	URL                       string    `json:"url,omitempty"`
	ForksURL                  string    `json:"forks_url,omitempty"`
	KeysURL                   string    `json:"keys_url,omitempty"`
	CollaboratorsURL          string    `json:"collaborators_url,omitempty"`
	TeamsURL                  string    `json:"teams_url,omitempty"`
	HooksURL                  string    `json:"hooks_url,omitempty"`
	IssueEventsURL            string    `json:"issue_events_url,omitempty"`
	EventsURL                 string    `json:"events_url,omitempty"`
	AssigneesURL              string    `json:"assignees_url,omitempty"`
	BranchesURL               string    `json:"branches_url,omitempty"`
	TagsURL                   string    `json:"tags_url,omitempty"`
	BlobsURL                  string    `json:"blobs_url,omitempty"`
	GitTagsURL                string    `json:"git_tags_url,omitempty"`
	GitRefsURL                string    `json:"git_refs_url,omitempty"`
	TreesURL                  string    `json:"trees_url,omitempty"`
	StatusesURL               string    `json:"statuses_url,omitempty"`
	LanguagesURL              string    `json:"languages_url,omitempty"`
	StargazersURL             string    `json:"stargazers_url,omitempty"`
	ContributorsURL           string    `json:"contributors_url,omitempty"`
	SubscribersURL            string    `json:"subscribers_url,omitempty"`
	SubscriptionURL           string    `json:"subscription_url,omitempty"`
	CommitsURL                string    `json:"commits_url,omitempty"`
	GitCommitsURL             string    `json:"git_commits_url,omitempty"`
	CommentsURL               string    `json:"comments_url,omitempty"`
	IssueCommentURL           string    `json:"issue_comment_url,omitempty"`
	ContentsURL               string    `json:"contents_url,omitempty"`
	CompareURL                string    `json:"compare_url,omitempty"`
	MergesURL                 string    `json:"merges_url,omitempty"`
	ArchiveURL                string    `json:"archive_url,omitempty"`
	DownloadsURL              string    `json:"downloads_url,omitempty"`
	IssuesURL                 string    `json:"issues_url,omitempty"`
	PullsURL                  string    `json:"pulls_url,omitempty"`
	MilestonesURL             string    `json:"milestones_url,omitempty"`
	NotificationsURL          string    `json:"notifications_url,omitempty"`
	LabelsURL                 string    `json:"labels_url,omitempty"`
	ReleasesURL               string    `json:"releases_url,omitempty"`
	DeploymentsURL            string    `json:"deployments_url,omitempty"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
	PushedAt                  time.Time `json:"pushed_at,omitempty"`
	GitURL                    string    `json:"git_url,omitempty"`
	SSHURL                    string    `json:"ssh_url,omitempty"`
	CloneURL                  string    `json:"clone_url,omitempty"`
	SvnURL                    string    `json:"svn_url,omitempty"`
	Homepage                  any       `json:"homepage,omitempty"`
	Size                      int       `json:"size,omitempty"`
	StargazersCount           int       `json:"stargazers_count,omitempty"`
	WatchersCount             int       `json:"watchers_count,omitempty"`
	Language                  string    `json:"language,omitempty"`
	HasIssues                 bool      `json:"has_issues,omitempty"`
	HasProjects               bool      `json:"has_projects,omitempty"`
	HasDownloads              bool      `json:"has_downloads,omitempty"`
	HasWiki                   bool      `json:"has_wiki,omitempty"`
	HasPages                  bool      `json:"has_pages,omitempty"`
	HasDiscussions            bool      `json:"has_discussions,omitempty"`
	ForksCount                int       `json:"forks_count,omitempty"`
	MirrorURL                 any       `json:"mirror_url,omitempty"`
	Archived                  bool      `json:"archived,omitempty"`
	Disabled                  bool      `json:"disabled,omitempty"`
	OpenIssuesCount           int       `json:"open_issues_count,omitempty"`
	License                   any       `json:"license,omitempty"`
	AllowForking              bool      `json:"allow_forking,omitempty"`
	IsTemplate                bool      `json:"is_template,omitempty"`
	WebCommitSignoffRequired  bool      `json:"web_commit_signoff_required,omitempty"`
	Topics                    []any     `json:"topics,omitempty"`
	Visibility                string    `json:"visibility,omitempty"`
	Forks                     int       `json:"forks,omitempty"`
	OpenIssues                int       `json:"open_issues,omitempty"`
	Watchers                  int       `json:"watchers,omitempty"`
	DefaultBranch             string    `json:"default_branch,omitempty"`
	AllowSquashMerge          bool      `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit          bool      `json:"allow_merge_commit,omitempty"`
	AllowRebaseMerge          bool      `json:"allow_rebase_merge,omitempty"`
	AllowAutoMerge            bool      `json:"allow_auto_merge,omitempty"`
	DeleteBranchOnMerge       bool      `json:"delete_branch_on_merge,omitempty"`
	AllowUpdateBranch         bool      `json:"allow_update_branch,omitempty"`
	UseSquashPrTitleAsDefault bool      `json:"use_squash_pr_title_as_default,omitempty"`
	SquashMergeCommitMessage  string    `json:"squash_merge_commit_message,omitempty"`
	SquashMergeCommitTitle    string    `json:"squash_merge_commit_title,omitempty"`
	MergeCommitMessage        string    `json:"merge_commit_message,omitempty"`
	MergeCommitTitle          string    `json:"merge_commit_title,omitempty"`
}
type Head struct {
	Label string `json:"label,omitempty"`
	Ref   string `json:"ref,omitempty"`
	Sha   string `json:"sha,omitempty"`
	User  User   `json:"user,omitempty"`
	Repo  Repo   `json:"repo,omitempty"`
}
type Base struct {
	Label string `json:"label,omitempty"`
	Ref   string `json:"ref,omitempty"`
	Sha   string `json:"sha,omitempty"`
	User  User   `json:"user,omitempty"`
	Repo  Repo   `json:"repo,omitempty"`
}
type Self struct {
	Href string `json:"href,omitempty"`
}
type HTML struct {
	Href string `json:"href,omitempty"`
}
type Issue struct {
	Href string `json:"href,omitempty"`
}
type Comments struct {
	Href string `json:"href,omitempty"`
}
type ReviewComments struct {
	Href string `json:"href,omitempty"`
}
type ReviewComment struct {
	Href string `json:"href,omitempty"`
}
type Commits struct {
	Href string `json:"href,omitempty"`
}
type Statuses struct {
	Href string `json:"href,omitempty"`
}
type Links struct {
	Self           Self           `json:"self,omitempty"`
	HTML           HTML           `json:"html,omitempty"`
	Issue          Issue          `json:"issue,omitempty"`
	Comments       Comments       `json:"comments,omitempty"`
	ReviewComments ReviewComments `json:"review_comments,omitempty"`
	ReviewComment  ReviewComment  `json:"review_comment,omitempty"`
	Commits        Commits        `json:"commits,omitempty"`
	Statuses       Statuses       `json:"statuses,omitempty"`
}
type PullRequest struct {
	URL                 string    `json:"url,omitempty"`
	ID                  int       `json:"id,omitempty"`
	NodeID              string    `json:"node_id,omitempty"`
	HTMLURL             string    `json:"html_url,omitempty"`
	DiffURL             string    `json:"diff_url,omitempty"`
	PatchURL            string    `json:"patch_url,omitempty"`
	IssueURL            string    `json:"issue_url,omitempty"`
	Number              int       `json:"number,omitempty"`
	State               string    `json:"state,omitempty"`
	Locked              bool      `json:"locked,omitempty"`
	Title               string    `json:"title,omitempty"`
	User                User      `json:"user,omitempty"`
	Body                any       `json:"body,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	ClosedAt            any       `json:"closed_at,omitempty"`
	MergedAt            any       `json:"merged_at,omitempty"`
	MergeCommitSha      any       `json:"merge_commit_sha,omitempty"`
	Assignee            any       `json:"assignee,omitempty"`
	Assignees           []any     `json:"assignees,omitempty"`
	RequestedReviewers  []any     `json:"requested_reviewers,omitempty"`
	RequestedTeams      []any     `json:"requested_teams,omitempty"`
	Labels              []any     `json:"labels,omitempty"`
	Milestone           any       `json:"milestone,omitempty"`
	Draft               bool      `json:"draft,omitempty"`
	CommitsURL          string    `json:"commits_url,omitempty"`
	ReviewCommentsURL   string    `json:"review_comments_url,omitempty"`
	ReviewCommentURL    string    `json:"review_comment_url,omitempty"`
	CommentsURL         string    `json:"comments_url,omitempty"`
	StatusesURL         string    `json:"statuses_url,omitempty"`
	Head                Head      `json:"head,omitempty"`
	Base                Base      `json:"base,omitempty"`
	Links               Links     `json:"_links,omitempty"`
	AuthorAssociation   string    `json:"author_association,omitempty"`
	AutoMerge           any       `json:"auto_merge,omitempty"`
	ActiveLockReason    any       `json:"active_lock_reason,omitempty"`
	Merged              bool      `json:"merged,omitempty"`
	Mergeable           any       `json:"mergeable,omitempty"`
	Rebaseable          any       `json:"rebaseable,omitempty"`
	MergeableState      string    `json:"mergeable_state,omitempty"`
	MergedBy            any       `json:"merged_by,omitempty"`
	Comments            int       `json:"comments,omitempty"`
	ReviewComments      int       `json:"review_comments,omitempty"`
	MaintainerCanModify bool      `json:"maintainer_can_modify,omitempty"`
	Commits             int       `json:"commits,omitempty"`
	Additions           int       `json:"additions,omitempty"`
	Deletions           int       `json:"deletions,omitempty"`
	ChangedFiles        int       `json:"changed_files,omitempty"`
}
type Repository struct {
	ID                       int       `json:"id,omitempty"`
	NodeID                   string    `json:"node_id,omitempty"`
	Name                     string    `json:"name,omitempty"`
	FullName                 string    `json:"full_name,omitempty"`
	Private                  bool      `json:"private,omitempty"`
	Owner                    Owner     `json:"owner,omitempty"`
	HTMLURL                  string    `json:"html_url,omitempty"`
	Description              string    `json:"description,omitempty"`
	Fork                     bool      `json:"fork,omitempty"`
	URL                      string    `json:"url,omitempty"`
	ForksURL                 string    `json:"forks_url,omitempty"`
	KeysURL                  string    `json:"keys_url,omitempty"`
	CollaboratorsURL         string    `json:"collaborators_url,omitempty"`
	TeamsURL                 string    `json:"teams_url,omitempty"`
	HooksURL                 string    `json:"hooks_url,omitempty"`
	IssueEventsURL           string    `json:"issue_events_url,omitempty"`
	EventsURL                string    `json:"events_url,omitempty"`
	AssigneesURL             string    `json:"assignees_url,omitempty"`
	BranchesURL              string    `json:"branches_url,omitempty"`
	TagsURL                  string    `json:"tags_url,omitempty"`
	BlobsURL                 string    `json:"blobs_url,omitempty"`
	GitTagsURL               string    `json:"git_tags_url,omitempty"`
	GitRefsURL               string    `json:"git_refs_url,omitempty"`
	TreesURL                 string    `json:"trees_url,omitempty"`
	StatusesURL              string    `json:"statuses_url,omitempty"`
	LanguagesURL             string    `json:"languages_url,omitempty"`
	StargazersURL            string    `json:"stargazers_url,omitempty"`
	ContributorsURL          string    `json:"contributors_url,omitempty"`
	SubscribersURL           string    `json:"subscribers_url,omitempty"`
	SubscriptionURL          string    `json:"subscription_url,omitempty"`
	CommitsURL               string    `json:"commits_url,omitempty"`
	GitCommitsURL            string    `json:"git_commits_url,omitempty"`
	CommentsURL              string    `json:"comments_url,omitempty"`
	IssueCommentURL          string    `json:"issue_comment_url,omitempty"`
	ContentsURL              string    `json:"contents_url,omitempty"`
	CompareURL               string    `json:"compare_url,omitempty"`
	MergesURL                string    `json:"merges_url,omitempty"`
	ArchiveURL               string    `json:"archive_url,omitempty"`
	DownloadsURL             string    `json:"downloads_url,omitempty"`
	IssuesURL                string    `json:"issues_url,omitempty"`
	PullsURL                 string    `json:"pulls_url,omitempty"`
	MilestonesURL            string    `json:"milestones_url,omitempty"`
	NotificationsURL         string    `json:"notifications_url,omitempty"`
	LabelsURL                string    `json:"labels_url,omitempty"`
	ReleasesURL              string    `json:"releases_url,omitempty"`
	DeploymentsURL           string    `json:"deployments_url,omitempty"`
	CreatedAt                time.Time `json:"created_at,omitempty"`
	UpdatedAt                time.Time `json:"updated_at,omitempty"`
	PushedAt                 time.Time `json:"pushed_at,omitempty"`
	GitURL                   string    `json:"git_url,omitempty"`
	SSHURL                   string    `json:"ssh_url,omitempty"`
	CloneURL                 string    `json:"clone_url,omitempty"`
	SvnURL                   string    `json:"svn_url,omitempty"`
	Homepage                 any       `json:"homepage,omitempty"`
	Size                     int       `json:"size,omitempty"`
	StargazersCount          int       `json:"stargazers_count,omitempty"`
	WatchersCount            int       `json:"watchers_count,omitempty"`
	Language                 string    `json:"language,omitempty"`
	HasIssues                bool      `json:"has_issues,omitempty"`
	HasProjects              bool      `json:"has_projects,omitempty"`
	HasDownloads             bool      `json:"has_downloads,omitempty"`
	HasWiki                  bool      `json:"has_wiki,omitempty"`
	HasPages                 bool      `json:"has_pages,omitempty"`
	HasDiscussions           bool      `json:"has_discussions,omitempty"`
	ForksCount               int       `json:"forks_count,omitempty"`
	MirrorURL                any       `json:"mirror_url,omitempty"`
	Archived                 bool      `json:"archived,omitempty"`
	Disabled                 bool      `json:"disabled,omitempty"`
	OpenIssuesCount          int       `json:"open_issues_count,omitempty"`
	License                  any       `json:"license,omitempty"`
	AllowForking             bool      `json:"allow_forking,omitempty"`
	IsTemplate               bool      `json:"is_template,omitempty"`
	WebCommitSignoffRequired bool      `json:"web_commit_signoff_required,omitempty"`
	Topics                   []any     `json:"topics,omitempty"`
	Visibility               string    `json:"visibility,omitempty"`
	Forks                    int       `json:"forks,omitempty"`
	OpenIssues               int       `json:"open_issues,omitempty"`
	Watchers                 int       `json:"watchers,omitempty"`
	DefaultBranch            string    `json:"default_branch,omitempty"`
}
type Sender struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}

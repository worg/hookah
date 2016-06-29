package webhooks

type (
	// GitHub hook structure
	GitHub struct {
		Head
		Commits    []gitHubCommit `json:"commits"`
		Compare    string         `json:"compare"`
		Created    bool           `json:"created"`
		Deleted    bool           `json:"deleted"`
		Forced     bool           `json:"forced"`
		HeadCommit gitHubCommit   `json:"head_commit"`
		Pusher     User           `json:"pusher"`
		Repository gitHubRepo     `json:"repository"`
	}

	gitHubCommit struct {
		Commit
		Added     []string      `json:"added"`
		Committer User          `json:"committer"`
		Distinct  bool          `json:"distinct"`
		Modified  []string      `json:"modified"`
		Removed   []interface{} `json:"removed"`
	}

	gitHubRepo struct {
		Repo
		CreatedAt    int    `json:"created_at"`
		Fork         bool   `json:"fork"`
		Forks        int    `json:"forks"`
		HasDownloads bool   `json:"has_downloads"`
		HasIssues    bool   `json:"has_issues"`
		HasWiki      bool   `json:"has_wiki"`
		ID           int    `json:"id"`
		Language     string `json:"language"`
		MasterBranch string `json:"master_branch"`
		OpenIssues   int    `json:"open_issues"`
		Owner        User   `json:"owner"`
		Private      bool   `json:"private"`
		PushedAt     int    `json:"pushed_at"`
		Size         int    `json:"size"`
		Stargazers   int    `json:"stargazers"`
		Watchers     int    `json:"watchers"`
	}
)

// Hook returns a CommonHook structure
// to ease handing of basic data
func (g GitHub) Hook() CommonHook {
	ch := CommonHook{
		Head:    g.Head,
		Repo:    g.Repository.Repo,
		Author:  g.Pusher,
		Commits: make([]Commit, len(g.Commits)),
	}

	for i, c := range g.Commits {
		ch.Commits[i] = c.Commit
	}

	return ch
}

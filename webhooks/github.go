package webhooks

type (
	// GitHub hook structure
	GitHub struct {
		hook
		Commits    []gitHubCommit `json:"commits"`
		Compare    string         `json:"compare"`
		Created    bool           `json:"created"`
		Deleted    bool           `json:"deleted"`
		Forced     bool           `json:"forced"`
		HeadCommit gitHubCommit   `json:"head_commit"`
		Pusher     user           `json:"pusher"`
		Repository gitHubRepo     `json:"repository"`
	}

	gitHubCommit struct {
		commit
		Added     []string      `json:"added"`
		Committer user          `json:"committer"`
		Distinct  bool          `json:"distinct"`
		ID        string        `json:"id"`
		Modified  []string      `json:"modified"`
		Removed   []interface{} `json:"removed"`
	}

	gitHubRepo struct {
		repo
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
		Owner        user   `json:"owner"`
		Private      bool   `json:"private"`
		PushedAt     int    `json:"pushed_at"`
		Size         int    `json:"size"`
		Stargazers   int    `json:"stargazers"`
		Watchers     int    `json:"watchers"`
	}
)

func (g *GitHub) Hook() CommonHook {
	ch := CommonHook{
		hook:    g.hook,
		Repo:    g.Repository.repo,
		Commits: make([]commit, len(g.Commits)),
	}

	for i, c := range g.Commits {
		ch.Commits[i] = c.commit
	}

	return ch
}

package webhooks

type (
	// GitLab hook structure
	GitLab struct {
		hook
		Commits           []commit   `json:"commits"`
		Kind              string     `json:"object_kind"`
		ProjectID         int        `json:"project_id"`
		Repository        gitLabRepo `json:"repository"`
		TotalCommitsCount int        `json:"total_commits_count"`
		UserEmail         string     `json:"user_email"`
		UserID            int        `json:"user_id"`
		UserName          string     `json:"user_name"`
		Attributes        gitLabAttr `json:"object_attributes"`
	}

	gitLabRepo struct {
		repo
		GitHttpURL      string `json:"git_http_url"`
		GitSshURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	}

	gitLabMerge struct {
		HttpURL         string `json:"http_url"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		SshURL          string `json:"ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	}

	gitLabAttr struct {
		Action          string      `json:"action"`
		AssigneeID      int         `json:"assignee_id"`
		AuthorID        int         `json:"author_id"`
		CreatedAt       string      `json:"created_at"`
		Description     string      `json:"description"`
		ID              int         `json:"id"`
		Iid             int         `json:"iid"`
		LastCommit      commit      `json:"last_commit"`
		MergeStatus     string      `json:"merge_status"`
		MilestoneID     string      `json:"milestone_id"`
		Source          gitLabMerge `json:"source"`
		SourceBranch    string      `json:"source_branch"`
		SourceProjectID int         `json:"source_project_id"`
		StCommits       string      `json:"st_commits"`
		StDiffs         string      `json:"st_diffs"`
		State           string      `json:"state"`
		Target          gitLabMerge `json:"target"`
		TargetBranch    string      `json:"target_branch"`
		TargetProjectID int         `json:"target_project_id"`
		Title           string      `json:"title"`
		UpdatedAt       string      `json:"updated_at"`
		URL             string      `json:"url"`
	}
)

func (g *GitLab) Hook() CommonHook {
	return CommonHook{
		hook:    g.hook,
		Repo:    g.Repository.repo,
		Commits: g.Commits,
	}
}
package webhooks

type (
	hook struct {
		After  string `json:"after"`
		Before string `json:"before"`
		Ref    string `json:"ref"`
	}

	commit struct {
		Author    user   `json:"author"`
		ID        string `json:"id"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
	}

	user struct {
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
		Name      string `json:"name"`
		Username  string `json:"username"`
	}

	repo struct {
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
		Name        string `json:"name"`
		URL         string `json:"url"`
	}

	CommonHook struct {
		hook
		Repo    repo
		Commits []commit
	}

	Context interface {
		Hook() CommonHook
	}
)

package webhooks

type (
	// Head is the base structure for hook payloads
	Head struct {
		After  string `json:"after"`
		Before string `json:"before"`
		Ref    string `json:"ref"`
	}

	// Commit holds basic commit information
	Commit struct {
		Author    User   `json:"author"`
		ID        string `json:"id"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
	}

	// User holds provider agnostic
	// basic user/commiter information
	User struct {
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
		Name      string `json:"name"`
		Username  string `json:"username"`
	}

	// Repo holds provider agnostic
	// git repository information
	Repo struct {
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
		Name        string `json:"name"`
		URL         string `json:"url"`
	}

	// CommonHook holds webhook payload information
	// regardless of the provider
	CommonHook struct {
		Head
		Author  User
		Repo    Repo
		Commits []Commit
	}

	// Context interface defines types
	// that provide a CommonHook struture
	Context interface {
		Hook() CommonHook
	}
)

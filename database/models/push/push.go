package push

type (
	Push struct {
		ID           string `json:"id" db:"id"`
		BeforeCommit string `json:"beforeCommit" db:"before_commit"`
		AfterCommit  string `json:"afterCommit" db:"after_commit"`
		PusherID     string `json:"pusherId" db:"pusher_id"`
		SenderID     string `json:"senderId" db:"sender_id"`
		DiffURL      string `json:"diffUrl" db:"diff_url"`
	}

	Pusher struct {
		ID    string `json:"id" db:"id"`
		Name  string `json:"name" db:"name"`
		Email string `json:"email" db:"email"`
	}

	Sender struct {
		ID        string `json:"id" db:"id"`
		Login     string `json:"login" db:"login"`
		HtmlURL   string `json:"htmlUrl" db:"html_url"`
		AvatarURL string `json:"avatarUrl" db:"avatar_url"`
		ApiURL    string `json:"apiUrl" db:"api_url"`
	}
)

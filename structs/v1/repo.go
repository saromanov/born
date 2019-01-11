package v1

// Repo defines object for repository
type Repo struct {
	ID       int64  `json:"id,omitempty"`
	UserID   int64  `json:"-"`
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Avatar   string `json:"avatar_url,omitempty"`
}

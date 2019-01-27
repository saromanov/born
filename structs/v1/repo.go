package v1

// Repo defines object for repository
type Repo struct {
	ID       int64  `json:"id,omitempty"`
	UserID   int64  `json:"-"`
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	CloneURL string `json:"clone_url,omitempty"`
}

// ContentFile provides definition of content file
// at the repo
type ContentFile struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

// GetContentRequest defines request for getting content of
//the file from repo
type GetContentRequest struct {
	u        *User
	repo     string
	owner    string
	fileName string
}

package v1

// Repo defines object for repository
type Repo struct {
	ID         int64  `json:"id,omitempty"`
	UserID     int64  `json:"-"`
	Owner      string `json:"owner"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	CloneURL   string `json:"clone_url,omitempty"`
	ArchiveURL string `json:"archive_url,omitempty"`
}

// ContentFile provides definition of content file
// at the repo
type ContentFile struct {
	ID      int64  `json:"id"`
	Content []byte `json:"content"`
}

// GetContentProvider defines request for getting content of
//the file from repo
type GetContentProvider struct {
	U        *User
	Repo     string
	Owner    string
	FileName string
}

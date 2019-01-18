package v1

// Build defines Born build
type Build struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Revision    int64     `json:"revision"`
	Status      string    `json:"status"`
	Event       string    `json:"event"`
	CreatedDate int64     `json:"created_date"`
	StartedDate int64     `json:"started_date"`
	Finished    int64     `json:"finished_date"`
	Branch      string    `json:"branch"`
	Steps       []Step    `json:"steps"`
	Services    []Service `json:"services,omitempty"`
	UserID      string    `json:"user_id"`
}

// Step defines for Build
type Step struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	Commands []string `json:"commands"`
	Pull     bool     `json:"pull"`
	Volumes  []Volume `json:"volumes"`
}

// Volume defines volume for Docker
type Volume struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Service defines service for build
type Service struct {
}

// BuildRequest defines struct at the POST request
// for making new build
type BuildRequest struct {
	Repo   string
	UserID string
}

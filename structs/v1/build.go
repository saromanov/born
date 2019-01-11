package v1

// Build defines Born build
type Build struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Revision int64     `json:"revision"`
	Steps    []Step    `json:"steps"`
	Services []Service `json:"services,omitempty"`
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

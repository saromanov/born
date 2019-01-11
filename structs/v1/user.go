package v1

// User defines registered user
type User struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Active bool   `json:"active"`
}

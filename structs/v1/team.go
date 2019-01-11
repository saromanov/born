package v1

// Team defines object for provider team
type Team struct {
	Name    string `json:"name,omitempty"`
	ID      string `json:"id"`
	Company string `json:"company,omitempty"`
}

package provider

import structs "github.com/saromanov/born/structs/v1"

// Provider defines interface for provider's platform
type Provider interface {
	Teams(*structs.User) ([]*structs.Team, error)
	Repo(*structs.User, int64) (*structs.Repo, error)
	GetContent(*structs.User, repo, path string)(*structs.ContentFile, error)
}

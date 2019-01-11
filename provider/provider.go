package provider

import structs "github.com/saromanov/born/structs/v1"

// Provider defines interface for provider's platform
type Provider interface {
	Teams(*structs.User) ([]*structs.Team, error)
	Repo(*structs.User, string) (*structs.Repo, error)
	Stats(*structs.User) error
}

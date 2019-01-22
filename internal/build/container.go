package build

import "github.com/fsouza/go-dockerclient"

// container provides handling of docker containers
type container struct {
	name   string
	client *docker.Client
}

func newContainer(name string) *container {
	return &container{
		name:   name,
		client: client,
	}
}

package build

import "github.com/fsouza/go-dockerclient"

// container provides handling of docker containers
type container struct {
	name   string
	client *docker.Client
}

// newContainer provides initialization of container struct
func newContainer(name string, c *docker.Client) *container {
	return &container{
		name:   name,
		client: c,
	}
}

// startContainer provides starting of container
func (c *container) startContainer() error {
	cont, err := c.client.CreateContainer(docker.CreateContainerOptions{
		Name: c.name,
		Config: &docker.Config{
			Image: "golang:latest",
		},
	})
	if err != nil {
		return err
	}

	err = c.client.StartContainer(cont.ID, &docker.HostConfig{})
	if err != nil {
		return err
	}

	return nil
}

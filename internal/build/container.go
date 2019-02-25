package build

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

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
			Hostname: "app",
		},
	})
	if err != nil {
		return err
	}
	fmt.Println("ID: ", cont.ID)
	err = c.client.StartContainer(cont.ID, &docker.HostConfig{})
	if err != nil {
		return err
	}

	return nil
}

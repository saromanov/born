package build

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
	"github.com/pkg/errors"
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
func (c *container) startContainer() (string, error) {
	cont, err := c.client.CreateContainer(docker.CreateContainerOptions{
		Name: c.name,
		Config: &docker.Config{
			Hostname: "app",
			Image:    c.name,
		},
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to start container")
	}
	fmt.Println("ID: ", cont.ID)
	err = c.client.StartContainer(cont.ID, &docker.HostConfig{})
	if err != nil {
		return "", errors.Wrap(err, "unable to start container")
	}

	err = c.client.RemoveContainer(docker.RemoveContainerOptions{
		ID: cont.ID,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to remove container")
	}

	return cont.ID, nil
}

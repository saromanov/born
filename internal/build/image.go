package build

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/fsouza/go-dockerclient"
	structs "github.com/saromanov/born/structs/v1"
)

const defaultEndpoint = "unix:///var/run/docker.sock"

var errImageNotDefined = errors.New("image is not defined")

// image defines structure for handling of Docker images
type image struct {
	client *docker.Client
	step   *structs.StepConfig
}

// newImage creates init for creating of docker images
func newImage(s *structs.StepConfig) (*image, error) {
	client, err := docker.NewClient(defaultEndpoint)
	if err != nil {
		return nil, err
	}
	return nil, &image{
		step:   s,
		client: client,
	}
}

// createDockerImage provides creating of the docker image from config
func createDockerImage(s *structs.StepConfig) error {
	var result string
	if s.Image == "" {
		return errImageNotDefined
	}
	result += fmt.Sprintf("FROM %s", s.Image)
	if len(s.Commands) > 0 {
		result += addCommands(s.Commands)
	}

	err := ioutil.WriteFile("/path1/Dockerfile", []byte(result), 0644)
	return err
}

func addCommands(c []string) string {
	var result string
	for i := 0; i < len(c); i++ {
		result += fmt.Sprintf("RUN %s", c[i])
	}
	return result
}

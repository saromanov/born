package build

import (
	"errors"
	"fmt"
	"io/ioutil"

	structs "github.com/saromanov/born/structs/v1"
)

var errImageNotDefined = errors.New("image is not defined")

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

package build

import (
	"fmt"
	"errors"
	"os"
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
}
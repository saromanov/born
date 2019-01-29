// Package build provides creating, starting and
// handling of builds
package build

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/fsouza/go-dockerclient"
	"github.com/saromanov/born/provider"
	structs "github.com/saromanov/born/structs/v1"
)

var errNoImage = errors.New("image is not defined")

// BuildStep provides definition for the build step
type BuildStep struct {
	Image    string
	Name     string
	Commands []string
	Parallel bool
}

// parseStep provides parsing of the step from the config
func parseStep(value interface{}) (BuildStep, error) {
	data := value.(map[interface{}]interface{})
	image, ok := data["image"]
	if !ok {
		return BuildStep{}, errNoImage
	}
	fmt.Println("IMAGE: ", image)
	s := BuildStep{
		Image: image.(string),
	}
	commands, ok := data["commands"]
	if ok {
		commandsOld := commands.([]interface{})
		commands := make([]string, len(commandsOld))
		for i := range commandsOld {
			commands[i] = commandsOld[i].(string)
		}
		s.Commands = commands
	}

	parallel, ok := data["parallel"]
	if ok {
		s.Parallel = parallel.(bool)
	}
	name, ok := data["name"]
	if ok {
		s.Name = name.(string)
	}
	return s, nil
}

// Build defines structure for build
type Build struct {
	P      provider.Provider
	User   *structs.User
	Repo   string
	mu     *sync.RWMutex
	images []string
}

// Create method provides creating of the build
func (b *Build) Create() error {
	b.mu = &sync.RWMutex{}
	c, err := b.getBornFile(b.Repo)
	if err != nil {
		return fmt.Errorf("unable to get born file: %v", err)
	}
	fmt.Println("CONFIG: ", c)
	client, err := newDockerClient()
	if err != nil {
		return err
	}

	b.images = make([]string, len(c.Steps))
	defer func(imgs []string) {
		for i := 0; i < len(imgs); i++ {
			client.RemoveImage(imgs[i])
		}
	}(b.images)
	for step, comm := range c.Steps {
		buildStep, err := parseStep(comm)
		if err != nil {
			continue
		}
		if buildStep.Parallel {
			go func(c *docker.Client, s string, bs BuildStep) {
				b.execuiteStep(c, s, bs)
			}(client, step, buildStep)
		} else {
			name, err := b.execuiteStep(client, step, buildStep)
			if err != nil {
				return err
			}
			b.images = append(b.images, name)
		}
	}
	return nil
}

// executeStep provides executing of the build step
func (b *Build) execuiteStep(client *docker.Client, step string, buildStep BuildStep) (string, error) {
	image := newImage(client)
	return image.createImage("1", step, buildStep)
}

// getBornFile provides getting of the .born.yml file
// from the repo. repo on format https://github.com/<owner>/<name>
func (b *Build) getBornFile(repo string) (*structs.Config, error) {
	res := strings.Split(repo, "/")
	resp, err := b.P.GetContent(&structs.GetContentProvider{
		Owner:    res[len(res)-2],
		Repo:     res[len(res)-1],
		FileName: ".born.yml",
	})
	if err != nil {
		return nil, err
	}
	return parseConfig(resp.Content)
}

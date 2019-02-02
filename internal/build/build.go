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

var (
	errNoImage           = errors.New("image is not defined")
	errRepoInvalidFormat = errors.New("repo url have invalid format")
)

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

	repo, err := b.getRepo(b.Repo)
	if err != nil {
		return fmt.Errorf("unable to get repo: %v", err)
	}
	fmt.Println("REPO: ", repo.ArchiveURL)
	client, err := newDockerClient()
	if err != nil {
		return err
	}

	//var wg sync.WaitGroup
	b.images = make([]string, len(c.Steps))
	for step, comm := range c.Steps {
		buildStep, err := parseStep(comm)
		if err != nil {
			fmt.Println("ERR: ", err)
			continue
		}
		/*if buildStep.Parallel {
			wg.Add(1)
			go func(c *docker.Client, s string, bs BuildStep) {
				name, err := b.execuiteStep(c, s, bs)
				if err == nil {
					b.images = append(b.images, name)
				}
				wg.Done()
			}(client, step, buildStep)
		} else {*/
		name, err := b.execuiteStep(client, step, buildStep)
		if err != nil {
			return err
		}
		b.images = append(b.images, name)
		//}
	}
	defer func(imgs []string) {
		for i := 0; i < len(imgs); i++ {
			client.RemoveImage(imgs[i])
		}
	}(b.images)

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
	owner, repoName, err := parseRepoURL(repo)
	if err != nil {
		return nil, err
	}
	resp, err := b.P.GetContent(&structs.GetContentProvider{
		Owner:    owner,
		Repo:     repoName,
		FileName: ".born.yml",
	})
	if err != nil {
		return nil, err
	}
	return parseConfig(resp.Content)
}

func (b *Build) getRepo(repo string) (*structs.Repo, error) {
	owner, repoName, err := parseRepoURL(repo)
	if err != nil {
		return nil, err
	}

	resp, err := b.P.Repo(nil, owner, repoName)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// parseRepoUrl returns owner and repo name
func parseRepoURL(repo string) (string, string, error) {
	res := strings.Split(repo, "/")
	if len(res) < 2 {
		return "", "", errRepoInvalidFormat
	}
	owner := res[len(res)-2]
	repoName := res[len(res)-1]
	return owner, repoName, nil
}

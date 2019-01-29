// Package build provides creating, starting and
// handling of builds
package build

import (
	"errors"
	"fmt"
	"strings"

	"github.com/saromanov/born/provider"
	structs "github.com/saromanov/born/structs/v1"
)

var errNoImage = errors.New("image is not defined")

// BuildStep provides definition for the build step
type BuildStep struct {
	Image    string
	Commands []string
}

func parseStep(value map[string]interface{}) (BuildStep, error) {
	image, ok := value["image"]
	if !ok {
		return BuildStep{}, errNoImage
	}

	return BuildStep{
		Image: image.(string),
	}, nil
}

// Build defines structure for build
type Build struct {
	P    provider.Provider
	User *structs.User
	Repo string
}

// Create method provides creating of the build
func (b *Build) Create() error {

	c, err := b.getBornFile(b.Repo)
	if err != nil {
		return fmt.Errorf("unable to get born file: %v", err)
	}
	fmt.Println("CONFIG: ", c)
	client, err := newDockerClient()
	if err != nil {
		return err
	}

	for step, comm := range c.Steps {
		image := newImage(client)
		name, err := image.createImage(b.User.ID, c.Steps[i])
		if err != nil {
			return err
		}
		fmt.Println(name)
	}
	for i := 0; i < len(c.Steps); i++ {
		/*image := newImage(client)
		name, err := image.createImage(b.User.ID, c.Steps[i])
		if err != nil {
			return err
		}
		container := newContainer(name, client)
		err = container.startContainer()
		if err != nil {
			continue
		}*/
	}
	return nil
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

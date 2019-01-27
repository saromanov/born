// Package build provides creating, starting and
// handling of builds
package build

import (
	"strings"

	"github.com/saromanov/born/provider"
	structs "github.com/saromanov/born/structs/v1"
)

// Build defines structure for build
type Build struct {
	p provider.Provider
}

// Create method provides creating of the build
func Create(u *structs.User, repo string) error {

	var c *structs.Config
	client, err := newDockerClient()
	if err != nil {
		return err
	}
	for i := 0; i < len(c.Steps); i++ {
		image := newImage(client)
		name, err := image.createImage(u.ID, c.Steps[i])
		if err != nil {
			return err
		}
		container := newContainer(name, client)
		err = container.startContainer()
		if err != nil {
			continue
		}
	}
	return nil
}

// getBornFile provides getting of the .born.yml file
// from the repo. repo on format https://github.com/<owner>/<name>
func (b *Build) getBornFile(repo string) error {
	res := strings.Split(repo, "/")
	_, err := b.p.Repo(nil, res[len(res)-2], res[len(res)-1])
	if err != nil {
		return err
	}
	return nil
}

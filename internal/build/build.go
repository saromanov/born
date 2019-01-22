// Package build provides creating, starting and
// handling of builds
package build

import structs "github.com/saromanov/born/structs/v1"

// Create method provides creating of the build
func Create(u *structs.User, repo string) error {

	var c *structs.Config
	image, err := newImage()
	if err != nil {
		return err
	}

	for i := 0; i < c.Steps; i++ {
		name, err := image.createImage(u.ID, c.Steps[i])
		if err != nil {
			return err
		}
		container, err := newContainer(name)
		if err != nil {
			return err
		}
	}
	return nil
}

// getBornFile provides getting of the .born.yml file
// from the repo
func getBornFile(repo string) error {

	return nil
}

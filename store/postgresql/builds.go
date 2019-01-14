package postgresql

import (
	"fmt"

	structs "github.com/saromanov/born/structs/v1"
)

// createBuild provides creating of teh new build and store
// it to the Postgresql
func (c *client) createBuild(b *structs.Build) error {
	if err := c.db.Create(b).Error; err != nil {
		return fmt.Errorf("unable to create build: %v", err)
	}

	return build
}

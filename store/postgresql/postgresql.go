package postgresql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/saromanov/born/store"
)

type client struct {
	db *gorm.DB
}

// New creates new init of the Postgresql
func New(opt *store.Options) store.Store {
	db, err := gorm.Open("postgres", fmt.Sprintf("user=%s password=%s"))
	if err != nil {
		panic("failed to connect database")
	}

	return client{
		db: db,
	}
}

// Close provides closing of db
func (c *client) Close() {
	c.db.Close()
}

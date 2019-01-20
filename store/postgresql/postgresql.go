package postgresql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saromanov/born/store"
)

type client struct {
	db *gorm.DB
}

// New creates new init of the Postgresql
func New(opt *store.Options) (store.Store, error) {
	fmt.Println(opt.Username, opt.Password)
	db, err := gorm.Open("postgres", fmt.Sprintf("user=%s password=%s", opt.Username, opt.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return client{
		db: db,
	}, nil
}

// Close provides closing of db
func (c *client) Close() {
	c.db.Close()
}

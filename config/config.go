package config

import (
	"sync"

	"github.com/cheesycoffee/go-template/config/internal/database"
	"github.com/cheesycoffee/go-template/config/internal/environment"
)

var (
	conf  *config
	mutex sync.Mutex
)

type (
	config struct {
		env environment.Environment
		db  database.Database
	}

	// Config application
	Config interface {
		DB() database.Database
		Env() environment.Environment
	}
)

// New Config
func New() Config {
	mutex.Lock()
	defer mutex.Unlock()
	if conf == nil {
		conf = new(config)
		conf.env = environment.NewEnvironment()
		conf.db = database.NewDatabase(database.Option{
			SQLReadDSN:    conf.env.SQLReadDSN,
			SQLWriteDSN:   conf.env.SQLWriteDSN,
			MongoReadURI:  conf.env.MongoReadURI,
			MongoWriteURI: conf.env.MongoWriteURI,
		})
		return conf
	}
	return nil
}

func (c *config) DB() database.Database {
	return c.db
}

func (c *config) Env() environment.Environment {
	return c.env
}

package config

import (
	"sync"

	"github.com/cheesycoffee/go-utilities/database"
)

var (
	conf  *config
	mutex sync.Mutex
)

type (
	config struct {
		env env
		db  database.Database
	}

	// Config application
	Config interface {
		DB() database.Database
		Env() env
	}
)

// New Config
func New() Config {
	mutex.Lock()
	defer mutex.Unlock()
	if conf == nil {
		conf = new(config)
		conf.env = newEnv()
		conf.db = database.NewDatabase(database.Option{
			SQLOption: database.SQLOption{
				ReadDSN:     conf.env.sqlReadDSN,
				WriteDSN:    conf.env.sqlWriteDSN,
				MaxIdleConn: conf.env.sqlMaxIdleConn,
				MaxOpenConn: conf.env.sqlMaxOpenConn,
				MaxLifeTime: conf.env.sqlMaxLifeTime,
			},
			MongoOption: database.MongoOption{
				ReadURI:  conf.env.mongoReadURI,
				WriteURI: conf.env.mongoWriteURI,
			},
			RedisOption: database.RedisOption{
				ReadDSN:  conf.env.redisReadDSN,
				WriteDSN: conf.env.redisWriteDSN,
			},
		})
		return conf
	}
	return nil
}

func (c *config) DB() database.Database {
	return c.db
}

func (c *config) Env() env {
	return c.env
}

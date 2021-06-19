package config

import (
	"time"
)

type env struct {
	Server     string `env:"SERVER"` // development, staging, production
	ServerPort string `env:"SERVER_PORT"`
	UseRest    bool   `env:"USE_REST"`
	UseGRPC    bool   `env:"USE_GRPC"`
	UseGraphQL bool   `env:"USE_GRAPHQL"`
	UseKafka   bool   `env:"USE_KAFKA"`

	sqlReadDSN     string        `env:"SQL_READ_DSN;required"`
	sqlWriteDSN    string        `env:"SQL_WRITE_DSN;required"`
	sqlMaxOpenConn int           `env:"SQL_MAX_OPEN_CONNECTION"`
	sqlMaxIdleConn int           `env:"SQL_MAX_IDLE_CONNECTION"`
	sqlMaxLifeTime time.Duration `env:"SQL_MAX_LIFETIME"`
	mongoReadURI   string        `env:"MONGO_READ_URI;required"`
	mongoWriteURI  string        `env:"MONGO_WRITE_URI;required"`
	redisReadDSN   string        `env:"REDIS_READ_DSN;required"`
	redisWriteDSN  string        `env:"REDIS_WRITE_DSN;required"`

	// Additional environment values
}

func newEnv() env {
	e := env{}
	library.MustParseEnv(&e)
	return e
}

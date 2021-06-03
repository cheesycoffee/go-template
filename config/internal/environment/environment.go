package environment

import "github.com/golangid/candi/candihelper"

// Environment server values
type Environment struct {
	Server     string `env:"SERVER"` // development, staging, production
	ServerPort string `env:"SERVER_PORT"`
	UseRest    bool   `env:"USE_REST"`
	UseGRPC    bool   `env:"USE_GRPC"`
	UseGraphQL bool   `env:"USE_GRAPHQL"`
	UseKafka   bool   `env:"USE_KAFKA"`

	SQLReadDSN    string `env:"SQL_READ_DSN"`
	SQLWriteDSN   string `env:"SQL_WRITE_DSN"`
	MongoReadURI  string `env:"MONGO_READ_URI"`
	MongoWriteURI string `env:"MONGO_WRITE_URI"`
}

// NewEnvironment server values
func NewEnvironment() Environment {
	env := Environment{}
	candihelper.MustParseEnv(&env)
	return env
}

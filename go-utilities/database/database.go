package database

type (
	database struct {
		gormSQL GormSQL
		mongodb MongoDB
		redisdb Redis
	}

	// Option dsn
	Option struct {
		SQLOption   SQLOption
		MongoOption MongoOption
		RedisOption RedisOption
	}

	// Database clients interface
	Database interface {
		GormSQL() GormSQL
		MongoDB() MongoDB
		Redis() Redis
	}
)

// NewDatabase clients
func NewDatabase(option Option) Database {
	dbase := new(database)
	dbase.gormSQL = newGormSQL(DBPostgreSQL, option.SQLOption)
	dbase.mongodb = newMongoDB(option.MongoOption)
	dbase.redisdb = newRedis(option.RedisOption)
	return dbase
}

func (db *database) GormSQL() GormSQL {
	return db.gormSQL
}

func (db *database) MongoDB() MongoDB {
	return db.mongodb
}

func (db *database) Redis() Redis {
	return db.redisdb
}

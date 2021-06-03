package database

// Database SQL type
const (
	DBMySQL      string = "mySQL"
	DBPostgreSQL string = "postgreSQL"
	DBSQLServer  string = "SQLServer"
	DBSQLite     string = "sqLite"
)

// Database Access type
const (
	DBAccessRead  string = "read"
	DBAccessWrite string = "write"
)

type (
	database struct {
		gormSQL GormSQL
		mongodb MongoDB
	}

	// Option dsn
	Option struct {
		SQLReadDSN, SQLWriteDSN     string
		MongoReadURI, MongoWriteURI string
	}

	// Database clients interface
	Database interface {
		GormSQL() GormSQL
		MongoDB() MongoDB
	}
)

// NewDatabase clients
func NewDatabase(option Option) Database {
	dbase := new(database)
	dbase.gormSQL = newGormSQL(DBPostgreSQL, SQLOption{
		ReadDSN:  option.SQLReadDSN,
		WriteDSN: option.SQLWriteDSN,
	})
	dbase.mongodb = newMongoDB(MongoOption{
		ReadURI:  option.MongoReadURI,
		WriteURI: option.MongoWriteURI,
	})
	return dbase
}

func (db *database) GormSQL() GormSQL {
	return db.gormSQL
}

func (db *database) MongoDB() MongoDB {
	return db.mongodb
}

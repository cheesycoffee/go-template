package database

import (
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database SQL type
const (
	DBMySQL      string = "mysql"
	DBPostgreSQL string = "postgres"
	DBSQLServer  string = "sqlserver"
	DBSQLite     string = "sqlite"
)

var openGormDB = gorm.Open

type (
	sqlInstance struct {
		read, write *gorm.DB
	}

	// SQLOption parameter
	SQLOption struct {
		ReadDSN, WriteDSN        string
		DebugMode                bool
		MaxOpenConn, MaxIdleConn int
		MaxLifeTime              time.Duration
	}

	// GormSQL database clients
	GormSQL interface {
		GetReadDB() *gorm.DB
		GetWriteDB() *gorm.DB
	}
)

func newGormSQL(dbType string, option SQLOption) GormSQL {
	sqlInst := new(sqlInstance)
	sqlInst.read = getGormSQL(dbType, false, option)
	sqlInst.write = getGormSQL(dbType, true, option)
	return sqlInst
}

func (s *sqlInstance) GetReadDB() *gorm.DB {
	return s.read
}

func (s *sqlInstance) GetWriteDB() *gorm.DB {
	return s.write
}

func getGormSQL(dbType string, isWriteAccess bool, option SQLOption) *gorm.DB {
	if strings.TrimSpace(option.ReadDSN) == "" || strings.TrimSpace(option.WriteDSN) == "" {
		return nil
	}

	var (
		dialector gorm.Dialector
		dsn       string = option.WriteDSN
	)

	gormConfig := &gorm.Config{}
	if !option.DebugMode {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	if !isWriteAccess {
		dsn = option.ReadDSN
		gormConfig.SkipDefaultTransaction = true
	}

	if strings.TrimSpace(dsn) == "" {
		return nil
	}

	switch dbType {
	case DBMySQL:
		dialector = mysql.Open(dsn)
	case DBSQLite:
		dialector = sqlite.Open(dsn)
	case DBSQLServer:
		dialector = sqlserver.Open(dsn)
	case DBPostgreSQL:
		dialector = postgres.Open(dsn)
	default:
		panic("Invalid SQL database type. The 'dbType' should on of the following : mySQL | postgreSQL | SQLServer | sqLite")
	}

	db, err := openGormDB(dialector, gormConfig)
	if err != nil {
		panic(err)
	}

	var (
		maxIdleConn = 10
		maxOpenConn = 100
		maxLifeTime = 30 * time.Minute
	)

	if option.MaxIdleConn > 0 {
		maxIdleConn = option.MaxIdleConn
	}

	if option.MaxOpenConn > 0 {
		maxOpenConn = option.MaxOpenConn
	}

	if option.MaxLifeTime > 0*time.Second {
		maxLifeTime = option.MaxLifeTime
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(maxLifeTime)

	return db
}

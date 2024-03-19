package connector

import (
	"database/sql"

	config "github.com/shing1211/go-lib/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB global variable to access gorm
var DB *gorm.DB

var sqlDB *sql.DB
var err error

// GetRdbmsDB - get a connection
func GetRdbmsDB() *gorm.DB {
	return DB
}

// InitRdbmsDB - function to initialize db
func InitRdbmsDB(config config.RDBMSConfig) *gorm.DB {
	var db = DB

	driver := config.RdbmsDbDriver
	username := config.RdbmsDbUser
	password := config.RdbmsDbPassword
	database := config.RdbmsDbName
	host := config.RdbmsDbHost
	port := config.RdbmsDbPort
	sslmode := config.RdbmsDbSslmode
	timeZone := config.RdbmsDbTimeZone
	maxIdleConns := config.RdbmsDbMaxIdleConns
	maxOpenConns := config.RdbmsDbMaxOpenConns
	connMaxLifetime := config.RdbmsDbConnMaxLifetime
	logLevel := config.RdbmsDbLogLevel

	switch driver {
	case "mysql":
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
		sqlDB, err = sql.Open(driver, dsn)
		if err != nil {
			log.WithError(err).Panic("panic code: 151")
		}
		sqlDB.SetMaxIdleConns(maxIdleConns)       // max number of connections in the idle connection pool
		sqlDB.SetMaxOpenConns(maxOpenConns)       // max number of open connections in the database
		sqlDB.SetConnMaxLifetime(connMaxLifetime) // max amount of time a connection may be reused

		db, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{
			Logger:         logger.Default.LogMode(logger.LogLevel(logLevel)),
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		})
		if err != nil {
			log.WithError(err).Panic("panic code: 152")
		}
		// Only for debugging
		if err == nil {
			log.Debug("Mysql/MariaDB connection successful!")
		}

	case "postgres":
		dsn := "host=" + host + " port=" + port + " user=" + username + " dbname=" + database + " password=" + password + " sslmode=" + sslmode + " TimeZone=" + timeZone
		sqlDB, err = sql.Open(driver, dsn)
		if err != nil {
			log.WithError(err).Panic("panic code: 153")
		}
		sqlDB.SetMaxIdleConns(maxIdleConns)       // max number of connections in the idle connection pool
		sqlDB.SetMaxOpenConns(maxOpenConns)       // max number of open connections in the database
		sqlDB.SetConnMaxLifetime(connMaxLifetime) // max amount of time a connection may be reused

		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
		})
		if err != nil {
			log.WithError(err).Panic("panic code: 154")
		}
		// Only for debugging
		if err == nil {
			log.Debug("DB connection successful!")
		}

	case "sqlite3":
		db, err = gorm.Open(sqlite.Open(database), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.WithError(err).Panic("panic code: 155")
		}
		// Only for debugging
		if err == nil {
			log.Debug("DB connection successful!")
		}

	default:
		log.Fatal("The driver " + driver + " is not implemented yet")
	}

	DB = db

	return DB
}

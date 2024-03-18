package config

import "time"

type RDBMSConfig struct {
	// for relational database configurations
	RdbmsDbDriver   string
	RdbmsDbUser     string
	RdbmsDbPassword string
	RdbmsDbName     string
	RdbmsDbHost     string
	RdbmsDbPort     string

	RdbmsDbSslmode  string
	RdbmsDbTimeZone string

	RdbmsDbMaxIdleConns    int
	RdbmsDbMaxOpenConns    int
	RdbmsDbConnMaxLifetime time.Duration
	RdbmsDbLogLevel        int
}

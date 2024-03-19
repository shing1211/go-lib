package config

import (
	"strconv"
	"time"

	oslib "github.com/shing1211/go-lib/os"
	stringlib "github.com/shing1211/go-lib/string"
	log "github.com/sirupsen/logrus"
)

// Configuration - Server, Db and Logger configuration variables
type Configuration struct {
	RabbitMQ RabbitMQConfig
	MongoDB  MongoConfig
	Server   ServerConfig
	Redis    RedisConfig
	Elastic  ElasticConfig
	Rdbms    RDBMSConfig
}

// Configuration - load the configurations from os.env
func Config() Configuration {
	var configuration Configuration
	getOSEnv := oslib.GetOSEnv

	// for elastic search configurations
	elasticHost := getOSEnv("ELASTIC_HOST", "elasticsearch")
	elasticPort := string(getOSEnv("ELASTIC_PORT", "9200"))
	elasticUser := getOSEnv("ELASTIC_USER", "elastic")
	elasticPwd := getOSEnv("ELASTIC_PWD", "changeme")

	// for mongo configurations
	mongoDBHost := getOSEnv("MONGODB_HOST", "mongo")
	mongoDBPort := string(getOSEnv("MONGODB_PORT", "27017"))
	mongoDBUser := getOSEnv("MONGODB_USER", "metavbot")
	mongoDBPwd := getOSEnv("MONGODB_PWD", "metavbot")
	MongoDBName := getOSEnv("MONGODB_DBNAME", "metavbot")

	// for rabbitmq configuration
	rabbitMQHost := getOSEnv("RABBITMQ_HOST", "rabbitmq")
	rabbitMQPort := string(getOSEnv("RABBITMQ_PORT", "5672"))
	rabbitMQUser := getOSEnv("RABBITMQ_USER", "guest")
	rabbitMQPwd := getOSEnv("RABBITMQ_PWD", "guest")

	// for Redis configurations
	redisHost := getOSEnv("REDIS_HOST", "redis")
	redisPort := string(getOSEnv("REDIS_PORT", "6379"))
	redisUser := getOSEnv("REDIS_USER", "")
	redisPwd := getOSEnv("REDIS_PWD", "")

	// for RDMBS configurations
	rdbmsDbDriver := getOSEnv("RDBMS_DRIVER", "mysql")
	rdbmsDbUser := getOSEnv("RDBMS_USER", "metavbot")
	rdbmsDbPassword := getOSEnv("RDBMS_PASSWORD", "metavbot")
	rdbmsDbName := getOSEnv("RDBMS_DBNAME", "metavbot")
	rdbmsDbHost := getOSEnv("RDBMS_HOST", "mariadb")
	rdbmsDbPort := string(getOSEnv("RDBMS_PORT", "3306"))
	rdbmsDbSslmode := getOSEnv("RDBMS_SSLMODE", "disable")
	rdbmsDbTimeZone := getOSEnv("RDBMS_TIMEZONE", "Asia/Hong_Kong")
	rdbmsDbMaxIdleConns := getOSEnv("RDBMS_MAXIDLECONNS", "10")
	rdbmsDbMaxOpenConns := getOSEnv("RDBMS_MAXOPENCONNS", "100")
	rdbmsDbConnMaxLifetime := getOSEnv("RDBMS_CONNMAXLIFETIME", "1h")
	rdbmsDbLogLevel := getOSEnv("RDBMS_LOGLEVEL", "4")

	serverPort := string(getOSEnv("GOLANG_API_SERVER_PORT", "8090"))
	serverMode := getOSEnv("GOLANG_API_SERVER_MODE", "development")
	serverLogLevel := getOSEnv("GOLANG_API_LOG_LEVEL", "6")

	mySigningKey := getOSEnv("MySigningKey", "Use_a_strong_and_long_random_key")
	JWTExpireTime, err := strconv.Atoi(getOSEnv("JWTExpireTime", "2"))
	if err != nil {
		log.WithError(err).Panic("panic code: 111")
	}

	hashPassMemory64, err := strconv.ParseUint((getOSEnv("HASHPASSMEMORY", "64")), 10, 64)
	if err != nil {
		log.WithError(err).Panic("panic code: 121")
	}

	hashPassIterations64, err := strconv.ParseUint((getOSEnv("HASHPASSITERATIONS", "2")), 10, 64)
	if err != nil {
		log.WithError(err).Panic("panic code: 122")
	}

	hashPassParallelism64, err := strconv.ParseUint((getOSEnv("HASHPASSPARALLELISM", "2")), 10, 64)
	if err != nil {
		log.WithError(err).Panic("panic code: 123")
	}

	hashPassSaltLength64, err := strconv.ParseUint((getOSEnv("HASHPASSSALTLENGTH", "16")), 10, 64)
	if err != nil {
		log.WithError(err).Panic("panic code: 124")
	}

	hashPassKeyLength64, err := strconv.ParseUint((getOSEnv("HASHPASSKEYLENGTH", "32")), 10, 64)
	if err != nil {
		log.WithError(err).Panic("panic code: 125")
	}

	hashPassMemory := uint32(hashPassMemory64)
	hashPassIterations := uint32(hashPassIterations64)
	hashPassParallelism := uint8(hashPassParallelism64)
	hashPassSaltLength := uint32(hashPassSaltLength64)
	hashPassKeyLength := uint32(hashPassKeyLength64)

	logLevel, err := stringlib.StringTouInt64(serverLogLevel)
	if err != nil {
		log.Warn(err)
	}
	configuration.Server.ServerLogLevel = logLevel

	configuration.Elastic.ElasticHost = elasticHost
	configuration.Elastic.ElasticPort = elasticPort
	configuration.Elastic.ElasticUser = elasticUser
	configuration.Elastic.ElasticPwd = elasticPwd

	configuration.RabbitMQ.RabbitMQHost = rabbitMQHost
	configuration.RabbitMQ.RabbitMQPort = rabbitMQPort
	configuration.RabbitMQ.RabbitMQUser = rabbitMQUser
	configuration.RabbitMQ.RabbitMQPwd = rabbitMQPwd

	configuration.MongoDB.MongoDBHost = mongoDBHost
	configuration.MongoDB.MongoDBPort = mongoDBPort
	configuration.MongoDB.MongoDBUser = mongoDBUser
	configuration.MongoDB.MongoDBPwd = mongoDBPwd
	configuration.MongoDB.MongoDBName = MongoDBName

	configuration.Redis.RedisHost = redisHost
	configuration.Redis.RedisPort = redisPort
	configuration.Redis.RedisUser = redisUser
	configuration.Redis.RedisPwd = redisPwd

	configuration.Rdbms.RdbmsDbDriver = rdbmsDbDriver
	configuration.Rdbms.RdbmsDbUser = rdbmsDbUser
	configuration.Rdbms.RdbmsDbPassword = rdbmsDbPassword
	configuration.Rdbms.RdbmsDbName = rdbmsDbName
	configuration.Rdbms.RdbmsDbHost = rdbmsDbHost
	configuration.Rdbms.RdbmsDbPort = rdbmsDbPort
	configuration.Rdbms.RdbmsDbSslmode = rdbmsDbSslmode
	configuration.Rdbms.RdbmsDbTimeZone = rdbmsDbTimeZone

	configuration.Server.ServerPort = serverPort
	configuration.Server.ServerMode = serverMode

	configuration.Rdbms.RdbmsDbMaxIdleConns, err = strconv.Atoi(rdbmsDbMaxIdleConns)
	if err != nil {
		log.WithError(err).Panic("panic code: 131")
	}

	configuration.Rdbms.RdbmsDbMaxOpenConns, err = strconv.Atoi(rdbmsDbMaxOpenConns)
	if err != nil {
		log.WithError(err).Panic("panic code: 132")
	}

	configuration.Rdbms.RdbmsDbConnMaxLifetime, err = time.ParseDuration(rdbmsDbConnMaxLifetime)
	if err != nil {
		log.WithError(err).Panic("panic code: 133")
	}

	configuration.Rdbms.RdbmsDbLogLevel, err = strconv.Atoi(rdbmsDbLogLevel)
	if err != nil {
		log.WithError(err).Panic("panic code: 134")
	}

	configuration.Server.ServerJWT.Key = mySigningKey
	configuration.Server.ServerJWT.Expire = JWTExpireTime

	configuration.Server.ServerHashPass.Memory = hashPassMemory
	configuration.Server.ServerHashPass.Iterations = hashPassIterations
	configuration.Server.ServerHashPass.Parallelism = hashPassParallelism
	configuration.Server.ServerHashPass.SaltLength = hashPassSaltLength
	configuration.Server.ServerHashPass.KeyLength = hashPassKeyLength

	return configuration
}

package connector

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/shing1211/go-lib/config"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var redisClient *redis.Client
var redisHost string

func GetRedisClient() (*redis.Client, error) {
	if redisClient == nil {
		if err := InitRedisConn; err != nil {
			log.Warn("Failed to get/initialize Redis client connection", err)
			return nil, errors.New("failed to initialize Redis client connection")
		}
	}
	return redisClient, nil
}

func InitRedisConn(config config.RedisConfig) error {
	host := config.RedisHost
	port := config.RedisPort
	//user := config.Redis.RedisUser
	pwd := config.RedisPwd
	redisHost = host + ":" + port
	log.Info("Connecting to Redis at: " + redisHost)
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: pwd,
		DB:       0,
	})
	if client == nil {
		log.Warn(err)
		return err
	}
	redisClient = client
	log.Info("Connected to redis at: " + redisHost)
	return nil
}

func PingRedisConn() error {
	log.Info("Redis healthcheck at: " + redisHost)
	if pong, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Warn(pong, err)
		return err
	}
	log.Info("Redis is healthy at: " + redisHost)
	return nil
}

func CloseRedisConn() error {
	log.Info("Closing Redis client connection at: " + redisHost)
	if redisClient == nil {
		return nil
	}
	if err := redisClient.Close(); err != nil {
		log.Warn(err)
		return err
	}
	log.Info("Closed Redis client connection at: " + redisHost)
	return nil
}

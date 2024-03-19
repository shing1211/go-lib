package connector

import (
	"context"
	"errors"
	"time"

	"github.com/shing1211/go-lib/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Client
var mongoDBUri string
var mongoDBHost string

const connectTimeout = 10

func GetMongoDBClient() (*mongo.Client, error) {
	if mongoDB == nil {
		if err := InitMongoDBConn; err != nil {
			log.Warn("Failed to get/initialize MongoDB client connection", err)
			return nil, errors.New("failed to initialize MongoDB client connection")
		}
	}
	return mongoDB, nil
}

func InitMongoDBConn(config config.MongoConfig) error {
	host := config.MongoDBHost
	port := config.MongoDBPort
	user := config.MongoDBUser
	password := config.MongoDBPwd
	dbName := config.MongoDBName

	log.Info("Connecting to MongoDB at: " + host + ":" + port)
	mongoDBHost = "mongodb://" + host + ":" + port + "/" + dbName
	mongoDBUri = "mongodb://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName

	clientOptions := options.Client().ApplyURI(mongoDBUri)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Warn(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Warn(err)
		return err
	}

	mongoDB = client
	log.Info("Connected to MongoDB at: " + host + ":" + port)
	return nil
}

func PingMongoConn(config config.MongoConfig) error {
	log.Info("MongoDB Healthcheck")
	if mongoDB == nil {
		if err := InitMongoDBConn(config); err != nil {
			log.Warn("Failed to get/initialize MongoDB client connection", err)
			return err
		}
	}
	if err := mongoDB.Ping(context.TODO(), nil); err != nil {
		log.Warn(err)
		return err
	}
	log.Info("MongoDB is healthy")
	return nil
}

func OpenCollection(config config.MongoConfig, collectionName string) (*mongo.Collection, error) {
	dbName := config.MongoDBName
	if mongoDB == nil {
		if err := InitMongoDBConn(config); err != nil {
			log.Warn("Failed to get/initialize MongoDB client connection", err)
			return nil, err
		}
	}
	return mongoDB.Database(dbName).Collection(collectionName), nil
}

func CloseMongoConn() *error {
	log.Info("Closing MongoDB client connection at: " + mongoDBHost)
	if mongoDB == nil {
		return nil
	}
	if err := mongoDB.Disconnect(context.TODO()); err != nil {
		log.Warn(err)
		return &err
	}
	log.Info("Closed MongoDB client connection at: " + mongoDBHost)
	return nil
}

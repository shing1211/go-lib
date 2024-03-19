package connector

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	config "github.com/shing1211/go-lib/config"
)

var RabbitMQConn *amqp.Connection
var rabbitMQUrl string
var rabbitMQHost string

func InitRabbitMQConn(config config.RabbitMQConfig) error {
	host := config.RabbitMQHost
	port := config.RabbitMQPort
	user := config.RabbitMQUser
	pwd := config.RabbitMQPwd

	rabbitMQHost = "amqp://" + host + ":" + port + "/"
	rabbitMQUrl = "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/"

	fmt.Println("Connecting to RabbitMQ at: ", host+":"+port)
	conn, err := amqp.Dial(rabbitMQUrl)

	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ: ", err)
		return err
	}

	RabbitMQConn = conn
	fmt.Println("Connected to Rabbitmq at: ", host+":"+port)
	return nil
}

func CloseRabbitMQConn() {
	fmt.Println("Closing RabbitMQ connection at: " + rabbitMQHost)
	if RabbitMQConn == nil {
		return
	}
	if err := RabbitMQConn.Close(); err != nil {
		fmt.Println("Failed to close RabbitMQ connection:", err)
		return
	}
	fmt.Println("Closed RabbitMQ connection at: " + rabbitMQHost)
}

package app

import (
	"go-rabbitmq/helper"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Ch *amqp.Channel
var Queue amqp.Queue

func InitRabbitMQ() {
	err := godotenv.Load()
	helper.FailOnError(err, "Error loading .env file")

	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	queueName := os.Getenv("QUEUE_NAME")

	Conn, err = amqp.Dial(rabbitMQURL)
	helper.FailOnError(err, "Failed to connect to RabbitMQ")

	Ch, err = Conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")

	Queue, err = Ch.QueueDeclare(
		queueName, // name of the queue from .env
		true,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	helper.FailOnError(err, "Failed to declare a queue")
}


func CloseRabbitMQ() {
	Ch.Close()
	Conn.Close()
}

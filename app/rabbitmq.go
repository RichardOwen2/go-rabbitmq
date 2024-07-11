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
	exchangeName := os.Getenv("EXCHANGE_NAME")
	queueName := os.Getenv("QUEUE_NAME")

	Conn, err = amqp.Dial(rabbitMQURL)
	helper.FailOnError(err, "Failed to connect to RabbitMQ")

	Ch, err = Conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")

	err = Ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	helper.FailOnError(err, "Failed to declare an exchange")

	Queue, err = Ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	helper.FailOnError(err, "Failed to declare a queue")

	routingKeys := []string{"create", "update"}

	for _, key := range routingKeys {
		err = Ch.QueueBind(
			queueName,    // queue name
			key,          // routing key
			exchangeName, // exchange
			false,
			nil,
		)
		helper.FailOnError(err, "Failed to bind a queue")
	}
}

func CloseRabbitMQ() {
	Ch.Close()
	Conn.Close()
}

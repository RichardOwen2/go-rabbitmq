package main

import (
	"go-rabbitmq/app"
	"go-rabbitmq/service"
)

func main() {
	// Initialize RabbitMQ
	app.InitRabbitMQ()
	defer app.CloseRabbitMQ()

	// Handle incoming messages
	service.HandleMessages()
}

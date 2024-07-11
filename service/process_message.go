package service

import (
	"log"

	"github.com/streadway/amqp"
)

func ProcessMessage(d *amqp.Delivery) {
	switch d.RoutingKey {
	case "create":
		log.Printf("Received message with create routing key: %s", d.RoutingKey)
	case "update":
		log.Printf("Received message with update routing key: %s", d.RoutingKey)
	default:
		log.Printf("Received message with unknown routing key: %s", d.RoutingKey)
	}
}

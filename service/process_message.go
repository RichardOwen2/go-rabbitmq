package service

import (
	"fmt"

	"github.com/streadway/amqp"
)

func ProcessMessage(d *amqp.Delivery) {
	fmt.Printf("Received a message: %s\n", d.Body)
}

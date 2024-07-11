package service

import (
	"go-rabbitmq/app"
	"log"
)

func HandleMessages() {
	msgs, err := app.Ch.Consume(
		app.Queue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	// Create a channel to signal the end of the program
	forever := make(chan bool)

	// Go routine to process messages
	go func() {
		for d := range msgs {
			ProcessMessage(&d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

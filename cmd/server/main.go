package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
	internal/pubsub
)

func main() {
	connectionString := "amqp://guest:guest@localhost:5672/"

	connection, err := amqp.Dial(connectionString)
	if err != nil {
		fmt.Printf("Internal error: %v\n", err)
		return
	}
	defer connection.Close()
	defer fmt.Println("\nSignal received, program shutting down")

	fmt.Println("Connection successful")

	signalChan := make(chan os.Signal, 1)
	PublishJSON
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}

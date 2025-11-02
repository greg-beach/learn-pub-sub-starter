package main

import (
	"fmt"
	"log"

	"github.com/greg-beach/learn-pub-sub-starter/internal/gamelogic"
	"github.com/greg-beach/learn-pub-sub-starter/internal/pubsub"
	"github.com/greg-beach/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connectionString := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(connectionString)
	if err != nil {
		fmt.Printf("Internal error: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Peril game server connect to RabbitMQ!")

	publishCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create channel: %v", err)
	}

	_, queue, err := pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilTopic,
		"game_logs",
		"game_logs.*",
		pubsub.SimpleQueueDurable,
	)
	if err != nil {
		log.Fatalf("could not subscribe to pause: %v", err)
	}
	fmt.Printf("Queue %v declared and bound!\n", queue.Name)

	gamelogic.PrintServerHelp()

	for {
		input := gamelogic.GetInput()
		if len(input) == 0 {
			continue
		}

		switch input[0] {
		case "pause":
			fmt.Println("sending pause message")

			err = pubsub.PublishJSON(publishCh, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})
			if err != nil {
				log.Printf("could not publish time: %v", err)
			}
			fmt.Println("pause message sent!")

		case "resume":
			fmt.Println("resuming game")

			err = pubsub.PublishJSON(publishCh, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: false})
			if err != nil {
				log.Printf("could not publish time: %v", err)
			}
			fmt.Println("resume message sent!")

		case "quit":
			fmt.Println("exiting game")
			return

		default:
			fmt.Println("unknown command")
			continue
		}

	}

}

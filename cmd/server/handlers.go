package main

import (
	"fmt"

	"github.com/greg-beach/learn-pub-sub-starter/internal/gamelogic"
	"github.com/greg-beach/learn-pub-sub-starter/internal/pubsub"
	"github.com/greg-beach/learn-pub-sub-starter/internal/routing"
)

func handlerLogs() func(gl routing.GameLog) pubsub.Acktype {
	return func(gl routing.GameLog) pubsub.Acktype {
		defer fmt.Println("> ")
		err := gamelogic.WriteLog(gl)
		if err != nil {
			return pubsub.NackDiscard
		}
		return pubsub.Ack
	}
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/xconnio/xconn-go"
)

const topicName = "io.xconn.image"

func main() {
	session, err := xconn.ConnectAnonymous(context.Background(), "rs://192.168.0.109:8082", "realm1")
	if err != nil {
		log.Fatal(err)
	}

	subscribeResp := session.Subscribe(topicName, func(event *xconn.Event) {
		imageBytes, err := event.ArgBytes(0)
		if err != nil {
			log.Println(err)
			return
		}

		if err = os.WriteFile("output.png", imageBytes, 0644); err != nil {
			log.Println(err)
		}
	}).Do()
	if subscribeResp.Err != nil {
		log.Fatal(subscribeResp.Err)
	}
	log.Println("Subscribed to topic", topicName)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

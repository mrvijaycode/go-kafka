package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "console-consumer-21564",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{"vijay"}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)

		if err != nil {
			fmt.Printf("Error")
		} else {
			fmt.Printf("Received: -> %s \n", string(msg.Value))
		}
	}

	consumer.Close()

}

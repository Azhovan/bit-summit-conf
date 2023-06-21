package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

var (
	topicName  = "product"
	pubsubName = "pubsub"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Kill, os.Interrupt)

	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < 1000; i++ {
		select {
		case <-ticker.C:
			{
				message := fmt.Sprintf("product-%d", i)
				err := client.PublishEvent(context.Background(), pubsubName, topicName, []byte(message))
				if err != nil {
					panic(err)
				}

				fmt.Printf("message:%s published\n", message)
			}
		case <-stop:
			fmt.Println("supplier closed")
		}
	}

}

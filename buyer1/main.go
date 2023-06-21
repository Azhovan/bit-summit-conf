package main

import (
	"context"
	"fmt"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"net/http"
)

var (
	topicName  = "product"
	pubsubName = "pubsub"
	address    = ":8083"
)

func main() {
	s := daprd.NewService(address)
	subscription := &common.Subscription{
		PubsubName: pubsubName,
		Topic:      topicName,
		Route:      fmt.Sprintf("/%s", topicName),
	}

	if err := s.AddTopicEventHandler(subscription, handler); err != nil {
		fmt.Println("buyer1: error while adding topic subscription")
		panic(err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		fmt.Println("buyer1: error while starting service")
		panic(err)
	}
}

func handler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	fmt.Printf("buyer1: %v received\n", e.Data)
	return false, nil
}

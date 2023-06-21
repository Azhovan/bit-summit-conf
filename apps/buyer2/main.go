package main

import (
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"math/rand"
	"net/http"
	"time"
)

var (
	topicName  = "product"
	pubsubName = "pubsub"
	address    = ":8084"

	targetServiceName = "market"
	targetMethodName  = "/offer"
)

func main() {
	s := daprd.NewService(address)
	subscription := &common.Subscription{
		PubsubName: pubsubName,
		Topic:      topicName,
		Route:      fmt.Sprintf("/%s", topicName),
	}

	if err := s.AddTopicEventHandler(subscription, handler); err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

type Offer struct {
	From    string
	Product string
	Val     int
}

func handler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	product := fmt.Sprintf("%v", e.Data)
	fmt.Printf("buyer2: %s received\n", product)

	client, err := dapr.NewClient()
	if err != nil {
		return false, err
	}

	offer := offerFor(product, 1, 100)
	res, err := client.InvokeMethodWithContent(ctx, targetServiceName, targetMethodName, "post", &dapr.DataContent{
		Data:        offer,
		ContentType: "text/plain",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("buyer1: offer %s sent. respos:%s \n", offer, string(res))
	return false, nil
}

func offerFor(product string, min, max int) []byte {
	rand.Seed(time.Now().UnixNano())
	randomNr := rand.Intn(max-min+1) + min

	offer := Offer{
		From:    "buyer2",
		Product: product,
		Val:     randomNr,
	}
	payload, err := json.Marshal(offer)
	if err != nil {
		panic(err)
	}

	return payload
}

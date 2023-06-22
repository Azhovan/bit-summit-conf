package main

import (
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"net/http"
)

var (
	address   = ":8086"
	sateStore = "state"
)

func main() {
	s := daprd.NewService(address)

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := s.AddServiceInvocationHandler("/offer", handler); err != nil {
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

func handler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	offer := &Offer{}
	err = json.Unmarshal(in.Data, offer)
	if err != nil {
		return nil, err
	}

	fmt.Printf("offer for: %s from: %s, val: %d\n", offer.Product, offer.From, offer.Val)

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	key := fmt.Sprintf("%s.%s", offer.Product, offer.From)
	err = client.SaveState(ctx, sateStore, key, []byte(string(offer.Val)), nil)

	return nil, err
}

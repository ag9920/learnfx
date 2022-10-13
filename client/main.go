package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"

	"github.com/ag9920/learnfx/kitex_gen/learn/fx/item"
	"github.com/ag9920/learnfx/kitex_gen/learn/fx/item/learnfxservice"
)

func main() {
	client, err := learnfxservice.NewClient("learn fx client", client.WithHostPorts("[::1]:8888"))
	if err != nil {
		log.Fatal(err)
	}

	// insert item
	req := &item.CreateItemReq{
		Item: &item.Item{
			ID:   9324359,
			Name: "test name",
			Desc: "test desc",
			VisibleUsers: []int64{
				1234,
			},
		},
	}
	resp, err := client.CreateItem(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Default().Printf("create item id=%v", resp.ID)

	// query visible item
	for {
		req := &item.FilterVisibleItemsReq{
			UserID: 1234,
			ItemIDs: []int64{
				9324359,
				23435423,
			},
		}
		resp, err := client.FilterVisibleItems(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}

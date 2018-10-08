package client

import (
	"github.com/TomorrowWu/learn-golang/crawler/engine"
	"github.com/TomorrowWu/learn-golang/crawler_distributed/config"
	"github.com/TomorrowWu/learn-golang/crawler_distributed/rpcsupport"

	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, e := rpcsupport.NewClient(host)
	if e != nil {
		return nil, e
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			e := client.Call(config.ItemSaverRpc, item, &result)
			if e != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, e)
			}
		}
	}()

	return out, nil
}

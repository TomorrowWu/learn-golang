package main

import (
	"fmt"
	"log"
	"net/rpc"

	"learn-golang/crawler/engine"
	"learn-golang/crawler/scheduler"
	"learn-golang/crawler/zhenai/parser"
	"learn-golang/crawler_distributed/config"
	itemsaver "learn-golang/crawler_distributed/persist/client"
	"learn-golang/crawler_distributed/rpcsupport"
	worker "learn-golang/crawler_distributed/worker/client"
)

func main() {
	port := fmt.Sprintf(":%d", config.ItemSaverPort)
	itemChan, err := itemsaver.ItemSaver(port)
	if err != nil {
		panic(err)
	}

	poll := createClientPoll([]string{"127.0.0.1"})

	processor := worker.CreateProcessor(poll)

	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueuedScheduler{},
		WorkerCount:    100,
		ItemChan:       itemChan,
		RequestProcess: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPoll(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}

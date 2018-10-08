package client

import (
	"net/rpc"

	"learn-golang/crawler/engine"
	"learn-golang/crawler_distributed/config"
	"learn-golang/crawler_distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}

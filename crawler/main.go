package main

import (
	"awesomeProject1/crawler/engine"
	"awesomeProject1/crawler/scheduler"
	"awesomeProject1/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	//e.Run(engine.Request{
	//	Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}

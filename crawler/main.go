package main

import (
	"github.com/golearn/crawler/engine"
	"github.com/golearn/crawler/scheduler"
	"github.com/golearn/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 100,
	}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParserCityList,
	// })
	// engine.Run(engine.Request{
	// 	Url:        "http://album.zhenai.com/u/1472268642",
	// 	ParserFunc: parser.ParserProfile,
	// })

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParserCityList,
	})
}

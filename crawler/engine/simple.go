package engine

import (
	"log"

	"github.com/golearn/crawler/fetcher"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		if r.Url == "" {
			continue
		}

		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item : %+v", item)
		}
	}
}

func Worker(r Request) (ParserResult, error) {
	log.Println("Fetching url :", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Println("Fetcher: error fetching url", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}

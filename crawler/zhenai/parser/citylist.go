package parser

import (
	"regexp"

	"github.com/golearn/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	// limit := 3

	result := engine.ParserResult{}
	for _, m := range matchs {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParserPeopleList,
			})
		// limit--
		// if limit < 1 {
		// 	break
		// }
	}
	return result
}

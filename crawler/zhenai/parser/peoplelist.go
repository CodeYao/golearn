package parser

import (
	"regexp"

	"github.com/golearn/crawler/engine"
)

const peopleListRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParserPeopleList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(peopleListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParserProfile(c, name)
				},
			})
	}
	return result
}

package parser

import (
	"regexp"
	"strconv"

	"github.com/golearn/crawler/engine"
	"github.com/golearn/crawler/model"
)

//var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-312fdcc4>([^<]+)</h1>`)
var genderRe = regexp.MustCompile(`"genderString":"([女|男]士)",`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var heigihtRe = regexp.MustCompile(`"heightString":"([\d]+)cm"`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([未婚|离异]+)</div>`)

func ParserProfile(contents []byte, name string) engine.ParserResult {
	var profile model.Profile

	profile.Age, _ = strconv.Atoi(getValue(contents, ageRe))
	profile.Name = name //getValue(contents, nameRe)
	profile.Gender = getValue(contents, genderRe)
	profile.Heigiht, _ = strconv.Atoi(getValue(contents, heigihtRe))
	profile.Weight, _ = strconv.Atoi(getValue(contents, weightRe))
	profile.Income = getValue(contents, incomeRe)
	profile.Marriage = getValue(contents, marriageRe)

	result := engine.ParserResult{}
	result.Items = append(result.Items, profile)
	result.Requests = append(
		result.Requests,
		engine.Request{
			Url:        "",
			ParserFunc: engine.NilParser,
		})
	return result
}

func getValue(contents []byte, re *regexp.Regexp) string {
	matchs := re.FindSubmatch(contents)
	if len(matchs) >= 2 {
		return string(matchs[1])
	} else {
		return ""
	}
}

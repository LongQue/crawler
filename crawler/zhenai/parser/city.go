package parser

import (
	"awesomeProject1/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//解析页面用户信息，提取出用户名
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			//将用户详情页加入待爬队列
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})

	}

	//解析城市信息
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}

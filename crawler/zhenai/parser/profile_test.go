package parser

import (
	"awesomeProject1/crawler/fetcher"
	"awesomeProject1/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch("http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "寂寞成影萌宝")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:       "寂寞成影萌宝",
		Gender:     "女",
		Age:        83,
		Height:     105,
		Weight:     137,
		Income:     "财务自由",
		Marriage:   "离异",
		Education:  "初中",
		Occupation: "金融",
		Hukou:      "南京市",
		Xinzuo:     "狮子座",
		House:      "无房",
		Car:        "无车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was%v", expected, profile)
	}

}

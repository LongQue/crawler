package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	printCityList(all)
}
func printCityList(contents []byte) {
	//<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/zunyi"class = "">遵义</a>
	re := regexp.MustCompile(`<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	matchs := re.FindAll(contents, -1)
	for _,m:=range matchs{
		fmt.Printf("%s/n",m)
	}
	fmt.Printf("Matches founc: %d\n",len(matchs))
}
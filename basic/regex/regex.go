package main

import (
	"fmt"
	"regexp"
)

const text = "My email is zhoufufeng@qq.com"
func main() {
	re:= regexp.MustCompile("zhoufufeng@qq.com")
	match := re.FindString(text)
	fmt.Println(match)
}

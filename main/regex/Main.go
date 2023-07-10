package main

import (
	"fmt"
	"regexp"
)

var WorkContentReg = regexp.MustCompile(`《.*?》`)
var WorkFilterReg = regexp.MustCompile(`[《》]`)

func main() {
	txt := ExtractWorkTxt("《蜘蛛侠3》")
	fmt.Println(txt)
}

func ExtractWorkTxt(s string) []string {
	result := make([]string, 0)
	for _, item := range WorkContentReg.FindAllString(s, -1) {
		item = WorkFilterReg.ReplaceAllString(item, "")
		result = append(result, item)
	}
	return result
}

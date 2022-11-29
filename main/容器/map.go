package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["11"] = 1
	m["22"] = 2
	m["33"] = 3
	m["44"] = 4
	m["55"] = 5

	//fmt.Println(m)

	for s := range m {
		fmt.Println(s)
	}

	_, flag := m["23"]

	if !flag {
		fmt.Println("不存在此键")
	}

	//delete(m, "1")
	//fmt.Println(m)
}

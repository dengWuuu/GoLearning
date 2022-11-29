package main

import "fmt"

func main() {
	num := []int{2, 3, 4}

	var sum = 0
	for _, i := range num {
		fmt.Print(i, " ")
		sum += i
	}
	fmt.Println(sum)

	var str = "123456"
	for i := range str {
		fmt.Println(i, "：", str[i])
	}

}

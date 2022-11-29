package main

import "fmt"

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	if i != 1 {
		println(1)
	} else if true {
		println(2)
	}

	var x = 2

	switch x {
	case 1:
		println(1)
	case 2:
		println(2)
	}

}

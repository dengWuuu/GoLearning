package main

import "fmt"

func main() {
	s := make([]int, 3)

	fmt.Println(s)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println("get", s)
	fmt.Println("len:", len(s))

	s = append(s, 12, 23, 345, 56, 452, 236)
	fmt.Println("append later:", s)

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:4]
	println("lh", l)
	l = s[2:]
	for _, i2 := range l {
		print(i2, " ")
	}

	n := map[string]int{"foo": 1}
	fmt.Println(n)

}

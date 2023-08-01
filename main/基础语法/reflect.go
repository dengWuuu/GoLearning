package main

import "reflect"

func main() {
	a := []int{1, 2}
	b := []int{1, 2}
	println(reflect.DeepEqual(a, b))
}

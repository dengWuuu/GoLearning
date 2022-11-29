package main

import "fmt"
import "unsafe"

var x = 65
var (
	y = 10
)

const MAX_VALUE = 0x3f

func main() {
	var a string = "abc"
	fmt.Println("hello, world", a, x)
	fmt.Println(&y)
	println(MAX_VALUE)

	println(unsafe.Sizeof(MAX_VALUE), unsafe.Sizeof(x))
}

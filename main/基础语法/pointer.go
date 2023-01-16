package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	var ptr *int
	i := 1
	ptr = &i

	//ptr = new(int)

	fmt.Println("前面", *ptr)
	zeroptr(ptr)
	fmt.Println("后面", *ptr)

	i = 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

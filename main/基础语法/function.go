package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func mutiVals(nums ...int) {
	fmt.Println(nums)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// 多值返回函数
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	fmt.Println(c)

	//多参数
	mutiVals(1, 2, 3, 4, 5)
	n := make([]int, 5)

	n = append(n, 1, 2, 3, 4, 5)

	mutiVals(n...)

	//闭包
	seq := intSeq()

	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
}

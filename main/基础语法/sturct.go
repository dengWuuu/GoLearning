package main

import "fmt"

type person struct {
	name string
	age  int
	ppp  P
}

func (p person) printAge() {
	fmt.Println(p.age)
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func testPtr(p *person) {
	p.age = 0
}

func test(p person) {
	p.age = 100
}

type P struct {
	a string
	b string
}

func main() {

	//var p = person{name: "test", age: 10}
	//fmt.Println(p)
	//testPtr(&p)
	//fmt.Println(p)
	//test(p)
	//fmt.Println(p)
	//fmt.Println("分割线下面为测试结构体方法---------")
	//p.printAge()
	//fmt.Println("下面为自动解引用-------------------------")
	//fmt.Println(&person{name: "Ann", age: 40})
	//
	//fmt.Println(newPerson("Jon"))
	//fmt.Println(*newPerson("Jon"))
	//
	//s := person{name: "Sean", age: 50}
	//fmt.Println(s.name)
	//
	//sp := &s
	//fmt.Println(sp.age)
	//
	//sp.age = 51
	//fmt.Println(sp.age)
	//
	//sp.age = 50
	//fmt.Println(sp.age)

	person := &person{
		name: "",
		age:  0,
	}
	fmt.Println(person.ppp)
}

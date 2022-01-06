package main

import "fmt"

// 类型别名
type MyInt int

type Person struct {
	name, city string
	age        int
}

func main() {
	// 1. 创建对象
	// 1.1 new
	p1 := new(Person)
	p1.name = "Alice"
	p1.age = 22
	// 1.2
	p2 := &Person{"clz", "123", 12}
	fmt.Println(p1)
	fmt.Println(p2)
}

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

func (h *Human) SayHi() {
	println("Hi")
}

func main() {
	s := &Student{Human{"clz", 18, "123231"}, "wust"}
	fmt.Printf("%+v\n", *s)
	s.SayHi()
}

package main

import (
	"fmt"
)

//Sayer 定义
type Sayer interface {
	Say()
}

//person 结构体
type person struct {
	Name string // 共有
	age  int    // 私有
}

//Person 构造函数
func Person(name string, age int) *person {
	return &person{Name: name, age: age}
}

//GetName 对象方法, 公有
func (p *person) GetName() string {
	return p.Name
}

//getAge 对象方法, 私有
func (p *person) getAge() int {
	return p.age
}

//say 实现接口
func (p *person) Say() {
	fmt.Println("My name is ", p.GetName())
}

func main() {
	p := Person("车亮召", 22)
	fmt.Println(p)
	fmt.Println(p.GetName())
	fmt.Println(p.getAge())
	var s Sayer
	s = p
	s.Say()
}

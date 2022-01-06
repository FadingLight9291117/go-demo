package main

import "fmt"

func main() {
	n0 := 1
	var nPtr *int = &n0
	fmt.Println(nPtr)
	fmt.Println(*nPtr)

	// 空指针
	var ptr0 *int = nil
	fmt.Println(ptr0)

	// new() 返回零值指针
	a := new(int)
	fmt.Println(*a)

	// make() 只适用于slice, map, chan的创建
	// 返回的就是三种类型本身，而不是指针类型
	// 因为这三种类型本身就是引用类型
	b := make(map[string]int, 10)
	b["test"] = 100
	fmt.Println(b)
}

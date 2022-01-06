package main

import (
	"fmt"
)

func main() {
	s := "asdads"
	for i, v := range s {
		fmt.Printf("%d -> %c\n", i, v)
	}
	fmt.Println(len(s))
	var m int32 = -(1 << 31)
	fmt.Println(m)
	x := '1'
	fmt.Println(foo(x))
	fmt.Println(len("1"))

	foo := make([]string, 0)
	foo = append(foo, "123")
	fmt.Println(foo)

	as := "a"
	fmt.Println(as[0])

	ss := "abc"
	fmt.Println(ss[0]== 'a')
	for _, v := range ss {
		fmt.Println(v)
	}
}

func foo(x rune) int {
	return int(x - '0')
}

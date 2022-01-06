package main

import "fmt"

//main array是值传递，slice是引用传递
func main() {
	arr := [...]int{1, 2, 3, 5, 1, 2, 6, 7, 7}

	s0 := arr[:2]
	s0[0] = 0
	fmt.Println(s0)
	fmt.Println(arr)

	s1 := arr[0:2:6] // [low:high:max]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s6 := []int{2, 3, 4}
	fmt.Println(s6)

	// make创建切片
	var s2 []int = make([]int, 2, 5) // | ptr | len | cap |
	s3 := make([]int, 2)
	fmt.Println(s2)
	fmt.Println(s3)

	// append追加
	s4 := append(s1, 10, 10, 10)
	fmt.Println(s4)

	// 切片操作
	fmt.Println(s4[0:2])

	s5 := make([]int, 3)
	copy(s5, s4)
	fmt.Println(s5)
}

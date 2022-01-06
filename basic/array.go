package main

import "fmt"

func main() {
	var arr0 [5]int = [5]int{1, 3, 4}
	var arr1 = [...]int{2, 4, 5, 6, 2}
	arr2 := [2][5]int{{1, 2}, {3, 4}}
	for i, v := range arr0 {
		fmt.Printf("arr0[%d] -> %d\n", i, v)
	}
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(len(arr0))
	fmt.Println(cap(arr0))

	arr3 := [...]int{1, 2, 3, 10: 0} // 使用索引号初始化
	fmt.Println(arr3)

}

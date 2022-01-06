package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	list := [...]string{"x", "a", "c"}
	str := strings.Join(list[:], "+")

	fmt.Println(str)
	s := "a b c"

	s1 := strings.Split(s, " ")
	fmt.Println(s1)

	path := "D:/test.txt"
	isPath, _ := filepath.Split(path)
	fmt.Println(isPath)
}

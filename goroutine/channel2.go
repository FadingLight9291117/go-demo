package main

import "fmt"

//单向通道
func foo1(out chan<- int) {
	out <- 10
}

func foo2(in <-chan int) {
	var1 := <-in
	fmt.Println(var1)
}

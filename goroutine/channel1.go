package main

import (
	"fmt"
	"time"
)

var g = make(chan int)

func routine1() {
	g <- 10
	fmt.Println("goroutine 1 done.")
}

func routine2() {
	var1 := <-g
	fmt.Println(var1)
	fmt.Println("goroutine 2 done.")
}

func main() {
	defer close(g)
	var ch1 chan int
	fmt.Println(ch1)
	ch2 := make(chan int, 10)
	fmt.Println(ch2)

	go routine2()
	go routine1()

	time.Sleep(time.Second)

}

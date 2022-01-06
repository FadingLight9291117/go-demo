package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go AFun(i)
	}
	time.Sleep(10 * time.Second)
}

func AFun(i int) {
	fmt.Printf("fun%d begin...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("fun%d end.\n", i)
}

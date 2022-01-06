package main


import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e1 := l.PushBack(1)
	l.PushBack(2)
	e3 := l.PushBack(3)

	l.InsertBefore(0, e1)
	l.InsertAfter(10, e3)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
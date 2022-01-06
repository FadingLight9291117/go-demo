package main

import (
	"fmt"
	"sync"
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) Do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {
	taskChan := make(chan task, 10)
	resultChan := make(chan int, 10)
	wait := &sync.WaitGroup{}
	go InitTasks(taskChan, resultChan, 100)
	go DistributeTask(taskChan, wait, resultChan)
	sum := ProcessResult(resultChan)
	fmt.Println(sum)
}

func InitTasks(taskChan chan<- task, resultChan chan int, n int) {
	stride := 10
	groups := n / stride
	rest := n % stride
	high := groups * stride
	for i := 0; i < groups; i++ {
		t := task{
			begin:  stride*i + 1,
			end:    (i + 1) * stride,
			result: resultChan,
		}
		taskChan <- t
	}
	if rest != 0 {
		t := task{
			begin:  high + 1,
			end:    n,
			result: resultChan,
		}
		taskChan <- t
	}
	close(taskChan)
}

func DistributeTask(taskChan <-chan task, wait *sync.WaitGroup, resultChan chan int) {
	for t := range taskChan {
		wait.Add(1)
		go ProcessTask(t, wait)
	}
	wait.Wait()
	close(resultChan)
}

func ProcessTask(t task, wait *sync.WaitGroup) {
	t.Do()
	wait.Done()
}

func ProcessResult(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		sum += r
	}
	return sum
}

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
	go ProcessTasks(taskChan, wait, resultChan)
	sum := ProcessResult(resultChan)
	fmt.Println(sum)

}

func InitTasks(taskChan chan<- task, resultChan chan int, n int) {
	groupN := n / 10
	rest := n % 10
	for i := 0; i < groupN; i++ {
		t := task{
			begin:  i*10 + 1,
			end:    (i + 1) * 10,
			result: resultChan,
		}
		taskChan <- t
	}
	if rest != 0 {
		begin := groupN*10 + 1
		end := begin + rest - 1
		t := task{
			begin:  begin,
			end:    end,
			result: resultChan,
		}
		taskChan <- t
	}
	close(taskChan)
}

func ProcessTasks(taskChan chan task, wait *sync.WaitGroup, resultChan chan int) {
	for t := range taskChan {
		wait.Add(1)
		go ProcessOneTask(t, wait)
	}
	wait.Wait()
	close(resultChan)
}

func ProcessOneTask(t task, wait *sync.WaitGroup) {
	t.Do()
	wait.Done()
}

func ProcessResult(resultChan <-chan int) int {
	sum := 0
	for i := range resultChan {
		sum += i
	}
	return sum
}

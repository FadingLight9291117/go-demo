package main

import (
	"fmt"
	"math/rand"
)

func InsertSort(data []int) {
	for i := 1; i < len(data); i++ {
		tmp := data[i]
		j := 0
		for j = i - 1; j >= 0 && tmp < data[j]; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = tmp
	}
}

func BubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func SelectSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func QuickSort(data []int) {
	if len(data) > 0 {
		pivot := data[0]
		i := 0
		j := len(data) - 1
		for true {
			for data[i] < pivot {
				i++
			}
			if i >= j {
				break
			}
			for data[j] > pivot {
				j--
			}
			if i >= j {
				break
			}
			data[i], data[j] = data[j], data[i]
		}
		data[i], data[0] = data[0], data[i]
		QuickSort(data[:i-1])
		QuickSort(data[i+1:])
	}
}

func main() {
	_data := [20]int{}
	for i := 0; i < 20; i++ {
		_data[i] = rand.Intn(20)
	}
	data := make([]int, len(_data))

	copy(data, _data[:])
	InsertSort(data)
	fmt.Println("InsertSort ->", data)

	copy(data, _data[:])
	BubbleSort(data)
	fmt.Println("BubbleSort ->", data)

	copy(data, _data[:])
	SelectSort(data)
	fmt.Println("SelectSort ->", data)

	copy(data, _data[:])
	QuickSort(data)
	fmt.Println("QuickSort  ->", data)
}

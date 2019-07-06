package main

import "fmt"

func SelectionSort(data []int) {
	size := len(data)
	var i, j, max int

	for i = 0; i < size-1; i++ {
		max = 0
		for j = 1; j < size-1-i; j++ {
			if data[j] > data[max] {
				max = j
			}
		}
		data[size-1-i], data[max] = data[max], data[size-1-i]
	}
}

func SelectionSort2(data []int) {
	size := len(data)
	var i, j, min int

	for i = 0; i < size-1; i++ {
		min = i
		for j = 1; j < size-1-i; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	SelectionSort(data)
	fmt.Println(data)
}

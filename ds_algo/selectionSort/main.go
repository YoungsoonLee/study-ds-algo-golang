package main

import "fmt"

func SelectionSort(data []int) {
	size := len(data)
	var i, j, min int

	for i = 0; i < size-1; i++ {
		min = i
		for j = i + 1; j < size; j++ {
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

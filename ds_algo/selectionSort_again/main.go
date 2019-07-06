package main

import "fmt"

func selecttionSort(data []int) {
	size := len(data)
	var min int

	for i := 0; i < size-1; i++ {
		min = i
		for j := i + 1; j < size; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

func main() {
	a := []int{3, 4, 1, 2, 10, 90, 8, 22}
	selecttionSort(a)
	fmt.Println(a)
}

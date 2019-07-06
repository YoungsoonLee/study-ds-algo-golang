package main

import "fmt"

func bubbleSort(data []int) []int {
	size := len(data)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func bubbleSortImprove(data []int) []int {
	size := len(data)
	swapped := 1

	for i := 0; i < size-1 && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = 1
			}
		}
	}
	return data
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	bubbleSortImprove(data)
	fmt.Println(data)

	data = []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	bubbleSort(data)
	fmt.Println(data)
}

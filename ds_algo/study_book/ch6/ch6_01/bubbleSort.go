package main

import "fmt"

func bubbleSort(data []int, comp func(int, int) bool) {
	size := len(data)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(data[j], data[j+1]) {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func bubbleSortImprove(data []int, comp func(int, int) bool) {
	size := len(data)
	swapped := 1

	for i := 0; i < size-1 && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(data[j], data[j+1]) {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = 1
			}
		}
	}
}

func more(v1 int, v2 int) bool {
	return v1 > v2
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	bubbleSortImprove(data, more)
	fmt.Println(data)

	bubbleSort(data, more)
	fmt.Println(data)
}

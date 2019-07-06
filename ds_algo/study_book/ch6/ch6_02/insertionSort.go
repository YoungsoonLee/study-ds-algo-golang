package main

import "fmt"

func insertionSort(data []int, comp func(int, int) bool) {
	size := len(data)
	var temp, i, j int

	for i = 1; i < size; i++ {
		temp = data[i]
		for j = i; j > 0 && comp(data[j-1], temp); j-- {
			data[j] = data[j-1]
		}
		data[j] = temp
	}
}

func more(v1 int, v2 int) bool {
	return v1 > v2
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	insertionSort(data, more)
	fmt.Println(data)
}

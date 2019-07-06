package main

import "fmt"

func insertionSort(data []int) {
	size := len(data)
	var temp, i, j int

	for i = 1; i < size; i++ {
		temp = data[i]
		for j = i - 1; j >= 0 && temp < data[j]; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}

}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	insertionSort(data)
	fmt.Println(data)
}

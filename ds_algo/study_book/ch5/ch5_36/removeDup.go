package main

import (
	"fmt"
	"sort"
)

func removeDup(data []int) int {
	j := 0
	size := len(data)
	if size == 0 {
		return 0
	}

	sort.Ints(data)

	for i := 1; i < size; i++ {
		if data[i] != data[j] {
			j++
			data[j] = data[i]
		}
	}

	return j + 1
}

func main() {
	a := []int{1, 2, 3, 3, 5, 6, 7, 8, 8}
	result := removeDup(a)
	fmt.Println(a, result)
}

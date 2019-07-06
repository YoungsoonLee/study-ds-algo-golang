package main

import (
	"fmt"
	"sort"
)

func findMedian(data []int) int {
	size := len(data)
	sort.Ints(data)
	return data[size/2]
}

func main() {
	a := []int{1, 2, 5, 6, 7, 8, 3, 4, 9}
	result := findMedian(a)
	fmt.Println(a, result)
}

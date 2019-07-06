package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(data []int) []int {
	size := len(data)
	if size == 1 {
		return data
	}
	mid := int(size / 2)

	left := data[:mid]
	right := data[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right)) //use extra memory

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			//result = append(result, left[0])
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// append leave left array
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	// append leave right array
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice), "\n")
}

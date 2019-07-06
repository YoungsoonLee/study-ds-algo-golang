package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generate a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(999))
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99) //- rand.Intn(999)
	}
	return slice
}

func mergeSort(items []int) []int {
	var size = len(items)

	if size == 1 {
		return items
	}

	// make a pivot
	middle := int(size / 2)
	//fmt.Println("middle: ", middle)

	/*
		var (
			left  = make([]int, middle)
			right = make([]int, size-middle)
		)


		for i := 0; i < size; i++ {
			if i < middle {
				left[i] = items[i]
			}else{
				right[i-middle] = items[i]
			}
		}
	*/
	left := items[:middle]
	right := items[middle:]

	//fmt.Printf("left: %v , right: %v \n", left, right)

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	// use extra memory
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// 한쪽만 남은거 처리
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	// 한쪽만 남은거 처리
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}

func main() {
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice), "\n")
}

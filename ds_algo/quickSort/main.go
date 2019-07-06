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

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1

	pivot := rand.Int() % len(data)
	//fmt.Println("pivot: ", pivot)

	data[pivot], data[right] = data[right], data[pivot]

	for i := range data {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left] // data i 와 left 를 바꾸고..
			left++                                    // left ++
		}
	}

	data[left], data[right] = data[right], data[left]

	//fmt.Println(data)

	quickSort(data[:left])
	quickSort(data[left+1:])
	return data
}

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quickSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

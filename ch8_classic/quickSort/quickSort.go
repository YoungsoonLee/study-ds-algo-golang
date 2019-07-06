package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99)
	}
	return slice
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a) // 하나의 피벗을 선정
	fmt.Println("pivot: ", pivot)

	a[pivot], a[right] = a[right], a[pivot] // 피벗을 맨 오른쪽으로 보냄.
	fmt.Println("before start: ", a)

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func qucickSorter(els []int, below, upper int) {
	if below < upper {
		pivot := dividePart(els, below, upper)
		qucickSorter(els, below, pivot-1)
		qucickSorter(els, pivot+1, upper)
	}
}

func dividePart(els []int, below, upper int) int {
	center := els[upper]
	i := below

	for j := below; j < upper; j++ {
		if els[j] <= center {
			els[i], els[j] = els[j], els[i]
			i++
		}
	}

	els[i], els[upper] = els[upper], els[i]
	return i
}

func main() {
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	//quickSort(slice)
	qucickSorter(slice, 0, len(slice)-1)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// To use a binary search algorithm,
// the list to be operated on must have already been sorted.
package main

import "fmt"

func binarySearch(needle int, stack []int) bool {
	low := 0
	high := len(stack) - 1

	for low <= high {
		median := (low + high) / 2
		if stack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(stack) || stack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(44, items))
}

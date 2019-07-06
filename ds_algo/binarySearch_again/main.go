package main

import "fmt"

// retun bool and index, if no data, redutn -1 of index
func binarySearch(data []int, key int) (bool, int) {
	low := 0
	high := len(data) - 1

	for low <= high {
		middele := (low + high) / 2
		fmt.Printf("low: %d, high: %d , middle: %d\n", low, high, middele)
		if data[middele] < key {
			low = middele + 1
		} else {
			high = middele - 1
		}
	}

	if low == len(data) || data[low] != key {
		return false, -1
	}

	return true, low
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(items, 45))
}

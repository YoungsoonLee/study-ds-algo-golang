package main

func BinarySearchRotateArray(data []int, key int) int {
	size := len(data)
	return BinarySearchRotateArrayUntil(data, 0, size-1, key)
}

func BinarySearchRotateArrayUntil(data []int, start int, end int, key int) int {
	if end < start {
		return -1
	}

	mid := (start + end) / 2

	if key == data[mid] {
		return mid
	}

	if data[mid] > data[start] {

	}
}

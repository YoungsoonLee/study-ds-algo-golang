package main

func MergeSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	tempArray := make([]int, size)
	mergeSrt(arr, tempArray, 0, size-1, comp)
}

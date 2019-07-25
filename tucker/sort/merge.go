package main

import "fmt"


func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])

	i, j := 0, 0
	rst := make([]int, 0)
	for i < len(left) || j < len(right) {
		if i >= len(left) { // 왼쪽 끝.
			rst = append(rst, right[j])
			j++
		} else if j >= len(right) {
			rst = append(rst, left[i])
			i++
		} else {
			if left[i] < right[j] {
				rst = append(rst, left[i])
				i++
			} else {
				rst = append(rst, right[j])
				j++
			}
		}
	}

	return rst

}

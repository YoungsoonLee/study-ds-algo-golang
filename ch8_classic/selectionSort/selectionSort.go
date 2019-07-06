package main

import "fmt"

func selection(nums []int) {
	l := len(nums)
	var min int

	for i := 0; i < l-1; i++ {
		min = i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		swap(&nums, i, min)
	}

	fmt.Println(nums)
}

func swap(a *[]int, i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func main() {
	n := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	selection(n)
}

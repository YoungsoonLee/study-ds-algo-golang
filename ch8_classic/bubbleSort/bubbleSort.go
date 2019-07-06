package main

import "fmt"

func bubbleSorter(nums []int) {
	var isSwapped = true // 멈출게 필요

	for isSwapped {
		isSwapped = false
		fmt.Println(isSwapped)
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				isSwapped = true
			}
			fmt.Println(nums)
		}

	}

	fmt.Println(nums)
}

func main() {
	n := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	bubbleSorter(n)
}

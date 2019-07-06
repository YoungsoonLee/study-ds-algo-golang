package main

import (
	"fmt"
	"sort"
)

func getMajority(data []int) (int, bool) {
	size := len(data)
	max := 0
	count := 0
	maxCount := 0

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}

		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}

	if maxCount > size/2 {
		return max, true
	}

	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

func getMajoritySort(data []int) (int, bool) {
	size := len(data)
	majIndex := size / 2
	sort.Ints(data) //sorted
	candidate := data[majIndex]
	count := 0

	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}

	fmt.Println(size / 2)

	if count > size/2 {
		fmt.Println(data[majIndex], count)
		return data[majIndex], true
	}

	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

func getMajority3(data []int) (int, bool) {
	majIndex := 0
	count := 1
	size := len(data)

	for i := 1; i < size; i++ {
		if data[majIndex] == data[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	fmt.Println("count: ", count, ", majIndex: ", majIndex)
	candidate := data[majIndex]
	fmt.Println("candidate: ", candidate)
	count = 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}

	if count > size/2 {
		fmt.Println(data[majIndex], count)
		return data[majIndex], true
	}

	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

func main() {
	a := []int{1, 2, 3, 3, 3, 3, 4}
	getMajority3(a)
}

package main

import "fmt"

func interpolationSearch(els []int, target int) (bool, int) {
	low, high := 0, len(els)-1
	var mid int

	for els[low] < target && els[high] > target {
		mid = low + ((target-els[low])*(high-low))/(els[high]-els[low])
		fmt.Println("mid: ", mid)

		if els[mid] < target {
			low = mid + 1
		} else if els[mid] > target {
			high = mid - 1
		} else {
			return true, mid
		}
	}

	if els[low] == target {
		return true, low
	} else if els[high] == target {
		return true, high
	} else {
		return false, -1
	}

	return false, -1
}

func main() {
	els := []int{2, 3, 5, 7, 9}
	t := 7

	f, i := interpolationSearch(els, t)
	fmt.Println(f, i)
}

package main

import "fmt"

func ShellSorter(els []int) {
	n := len(els)
	intervals := []int{1}
	k := 1

	for {
		interval := power(2, k) + 1
		if interval > n-1 {
			break
		}
		intervals = append([]int{interval}, intervals...)
		k++
	}

	fmt.Println("intervals: ", intervals)

	for _, interval := range intervals {
		for i := interval; i < n; i += interval {
			j := i
			for j > 0 {
				if els[j-interval] > els[j] {
					els[j-interval], els[j] = els[j], els[j-interval]
				}
				j = j - interval
			}
		}
	}
}

func power(exp, index int) int {
	power := 1
	for index > 0 {
		if index&1 != 0 {
			power *= exp
		}
		index >>= 1
		exp *= exp
	}
	return power
}

// main method
func main() {
	var elements []int
	elements = []int{34, 202, 13, 19, 6, 5, 1, 43, 506, 12, 20, 28, 17, 100, 25, 4, 5, 97, 1000, 27}
	ShellSorter(elements)
	fmt.Println(elements)
}

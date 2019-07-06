package main

import "fmt"

func maxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0

	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}

		fmt.Printf("maxEndingHere: %d , sofar: %d\n", maxEndingHere, maxSoFar)
	}

	return maxSoFar
}

func main() {
	data := []int{1, -2, 3, 4, 10, -4, 6, -14, 8, 2}
	fmt.Println(maxSubArraySum(data))
}

package main

import (
	"fmt"
	"math"
)

func largeContSun(arr []int) int {
	maxSum, currentSum := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		currentSum = int(math.Max(float64(currentSum+arr[i]), float64(arr[i])))
		maxSum = int(math.Max(float64(currentSum), float64(maxSum)))
	}

	return maxSum
}

func main() {
	a := []int{1, 2, -1, 3, 4, 10, 10, -10, -1}
	fmt.Println(largeContSun(a))
}

package main

import (
	"fmt"
)

func minAbsSumPair(data []int) {
	var sum int
	size := len(data)

	if size < 2 {
		fmt.Println("Invvalid inputs")
	}

	minFirst := 0
	minSecound := 1
	minSum := abs(data[0] + data[1])

	for i := 0; i < size; i++ {
		for r := i + 1; r < size; r++ {
			sum = abs(data[i] + data[r])
			if sum < minSum {
				minSum = sum
				minFirst = i
				minSecound = r
			}
		}
	}

	fmt.Println(" The two elements with minimum sum are : ", data[minFirst], " , ", data[minSecond])
}

package main

import "fmt"

func SumArray(data []int) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}

	fmt.Printf("Sum is %d\n", sum)
	return sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	SumArray(a)
}

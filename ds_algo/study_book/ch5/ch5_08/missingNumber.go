package main

import "fmt"

func missingNumber(data []int, dataRange int) int {
	size := len(data)
	xorSum := 0

	total := (dataRange + 1) * (dataRange / 2)
	fmt.Println(total)

	for i := 1; i < dataRange; i++ {
		xorSum ^= i
		fmt.Println("xorSum: ", xorSum)
	}
	fmt.Println("----")
	for i := 0; i < size; i++ {
		xorSum ^= data[i]
		fmt.Println("xorSum: ", xorSum)
	}
	fmt.Println("result: ", xorSum)
	return xorSum
}

func main() {
	a := []int{1, 2, 4, 5}
	missingNumber(a, 5)
}

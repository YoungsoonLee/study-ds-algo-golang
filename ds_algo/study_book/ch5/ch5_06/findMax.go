package main

import "fmt"

func findMax1(data []int) int {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1

	for i := 0; i < size; i++ {
		count = 1
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

	fmt.Println("max: ", max)
	return max
}

func main() {
	a := []int{1, 2, 3, 4, 4, 4, 5, 6, 7, 8, 8, 8, 8, 9}
	findMax1(a)

}

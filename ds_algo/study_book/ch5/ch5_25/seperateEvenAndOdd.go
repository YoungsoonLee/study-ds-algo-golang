package main

import "fmt"

func seperateEvenAndOdd(data []int) {
	size := len(data)
	left := 0
	right := size - 1

	for left < right {
		if data[left]%2 == 0 {
			left++
		} else if data[right]%2 == 1 {
			right--
		} else {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}
	fmt.Println(data)
}

func main() {
	a := []int{2, 3, 4, 5, 1, 7, 8, 9}
	seperateEvenAndOdd(a)
}

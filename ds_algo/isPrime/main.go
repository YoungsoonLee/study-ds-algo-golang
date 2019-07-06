package main

import (
	"fmt"
	"math"
)

func isPrime(data int) bool {
	if data < 2 {
		return false
	}

	for i := 2; i < data; i++ {
		if data%2 == 0 {
			return false
		}
	}

	return true
}

func isPrime2(data int) bool {
	if data < 2 {
		return false
	}

	sqrt := math.Sqrt(float64(data))
	fmt.Println("sqrt: ", int(sqrt))
	for i := 2; i <= int(sqrt); i++ {
		if data%2 == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(4))
	fmt.Println(isPrime2(3))
}

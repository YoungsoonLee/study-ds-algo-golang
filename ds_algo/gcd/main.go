package main

import (
	"fmt"
	"math"
)

func gcd(a, b int) int {
	i := int(math.Min(float64(a), float64(b)))

	for {
		if (a%i) == 0 || (b%i) == 0 {
			return i
		}
		i--
	}
}

func main() {
	fmt.Println(gcd(3, 6))
}

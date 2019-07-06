// 덧샘만을 이용해서 사칙연산 구현 하기

package main

import (
	"errors"
	"fmt"
	"math"
)

func negate(a int) int {
	neg := 0
	//d := a < 0 ? 1: -1

	d := 1
	if a > 0 {
		d = -1
	}

	for a != 0 {
		neg += d
		a += d
		/*
			fmt.Println("neg: ", neg)
			fmt.Println("a: ", a)
			fmt.Println("d: ", d)
		*/
	}
	return neg
}

func minus(a, b int) int {
	return a + negate(b)
}

func multi(a, b int) int {
	sum := 0
	for i := math.Abs(float64(b)); i > 0; i-- {
		sum += a
	}
	if b < 0 {
		sum = negate(sum)
	}
	return sum
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide zero")
	}
	absa := int(math.Abs(float64(a)))
	absb := int(math.Abs(float64(b)))

	x, p := 0, 0
	for p+absb <= absa {
		p += absb
		x++
	}

	if (a < 0 && b < 0) || (a > 0 && b > 0) {
		return x, nil
	} else {
		return negate(x), nil
	}

}

func main() {
	fmt.Println(minus(1, 3))
	fmt.Println(multi(3, -4))
	fmt.Println(divide(28, -4))
}

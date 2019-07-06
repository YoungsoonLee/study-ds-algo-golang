package main

import "fmt"

func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b
	}
}

func calc2(x int) int {
	a, b := 3, 5
	return a*x + b
}

func main() {
	f := calc()
	// f2 := calc2()

	//fmt.Println(f(1))
	//fmt.Println(calc2(1))

	//fmt.Println(f(2))
	//fmt.Println(calc2(2))

	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))
	fmt.Println(calc2(5))
}

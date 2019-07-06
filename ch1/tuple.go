package main

import "fmt"

func power(a int) (int, int) {
	return a * a, a * a * a
}

func main() {
	var squre int
	var cube int
	squre, cube = power(3)
	fmt.Println(squre, cube)
}

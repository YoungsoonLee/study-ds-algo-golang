package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 5; i++ {

		for j := 0; j < 5-i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < i*2+1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	// fmt.Println()
}

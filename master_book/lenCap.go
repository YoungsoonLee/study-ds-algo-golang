package main

import "fmt"

func printSlice(x []int) {
	for _, n := range x {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

func main() {
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	// 과거 Cap의 두배
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
}

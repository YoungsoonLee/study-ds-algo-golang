package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)

	integer := make([]int, 2)
	fmt.Println(integer)
	integer = nil
	fmt.Println(integer)

	anArray := [5]int{-1, -2, -3, -4, -5}
	refArray := anArray[:] // reference !!!!!!!!!!!!

	fmt.Println(anArray)
	fmt.Println(refArray)
	anArray[4] = -100
	fmt.Println(anArray)
	fmt.Println(refArray)

	s := make([]byte, 5)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}
	fmt.Println(twoD)
	fmt.Println()

	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i: ", i, "values: ", y)
		}
		fmt.Println()
	}
}

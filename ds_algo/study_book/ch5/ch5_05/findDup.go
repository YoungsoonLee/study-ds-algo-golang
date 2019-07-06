package main

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/set"
)

func printRepeatingSort(data []int) {
	size := len(data)
	sort.Ints(data)
	fmt.Println("Dup data are : ")

	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print(" ", data[i])
		}
	}

	fmt.Println()
}

func printRepeatingHash(data []int) {
	s := set.New()
	size := len(data)
	fmt.Println("Dup data are : ")
	for i := 0; i < size; i++ {
		if s.Has(data[i]) {
			fmt.Print(" ", data[i])
		} else {
			s.Insert(data[i])
		}
	}
	fmt.Println()
}

func printRepeatingCount(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)

	//set init
	for i := 0; i < size; i++ {
		count[i] = 0
	}

	fmt.Println("Dup data are : ")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Print(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()

}

func main() {
	a := []int{1, 2, 3, 4, 4, 4, 5, 6, 7, 8, 8, 8, 9}
	printRepeatingSort(a)

	printRepeatingHash(a)
}

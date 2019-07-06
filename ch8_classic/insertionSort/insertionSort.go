package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSequence(num int) []int {
	seq := make([]int, num, num)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < num; i++ {
		seq[i] = rand.Intn(999) - rand.Intn(999)
	}

	return seq
}

func InsertionSort(el []int) {
	n := len(el)
	//var i int

	for i := 1; i < n; i++ {
		var j int
		j = i
		for j > 0 {
			if el[j-1] > el[j] {
				el[j-1], el[j] = el[j], el[j-1]
			}
			j = j - 1 // !!!!!
		}
	}
}

func main() {

	var sequence []int
	sequence = randomSequence(24)
	fmt.Println("\n^^^^^^ Before Sorting ^^^ \n\n", sequence)
	InsertionSort(sequence)
	fmt.Println("\n--- After Sorting ---\n\n", sequence, "\n")
}

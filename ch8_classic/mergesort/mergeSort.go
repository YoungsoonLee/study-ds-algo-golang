package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createArray(num int) []int {
	array := make([]int, num, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		array[i] = rand.Intn(99) //- rand.Intn(99)
	}
	return array
}

func MergeSorter(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2

	//return JoinArray(MergeSorter(array[:mid]), MergeSorter(array[mid:]))
	return Merge(MergeSorter(array[:mid]), MergeSorter(array[mid:]))
}

//  Merge !!!
func JoinArray(left []int, right []int) []int {
	fmt.Println("left: ", left, ", right: ", right)

	num, i, j := len(left)+len(right), 0, 0
	res := make([]int, num, num)

	for k := 0; k < num; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			res[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			res[k] = left[i]
			i++
		} else if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
	}

	fmt.Println("res: ", res)
	return res

}

func Merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

// main method
func main() {

	var elements []int
	elements = createArray(10)
	fmt.Println("\n Before Sorting \n\n", elements)
	fmt.Println("\n-After Sorting\n\n", MergeSorter(elements), "\n")
}

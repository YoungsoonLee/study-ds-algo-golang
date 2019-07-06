package main

import "fmt"

func finder(arr1 []int, arr2 []int) int {

	sarr2 := map[int]int{}

	for _, num2 := range arr2 {
		sarr2[num2] = num2
	}

	//fmt.Println(sarr1)
	//fmt.Println(sarr2)

	for _, num := range arr1 {
		if _, ok := sarr2[num]; ok == false {
			return num
		}
	}

	return 0
}

func finderXor(arr1 []int, arr2 []int) int {
	result := 0
	//sum_arr := []int{}
	//fmt.Println(append(arr1, 10))
	for _, num := range arr2 {
		arr1 = append(arr1, num)
	}
	//fmt.Println(arr1)
	for _, num := range arr1 {
		result ^= num
	}

	return result
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{3, 7, 2, 1, 4, 6}

	fmt.Println(finder(arr1, arr2))
	fmt.Println(finderXor(arr1, arr2))
}

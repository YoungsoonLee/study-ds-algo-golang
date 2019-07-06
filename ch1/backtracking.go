package main

import "fmt"

func findElementsWithSum(arr [10]int, combi [19]int, size int, k int, addValue int, l, m int) int {
	var num int

	if addValue > k {
		return -1
	}

	if addValue == k {
		num = num + 1
		var p int
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combi[p]])
		}
		fmt.Println(" ")
	}

	var i int
	for i = l; i < size; i++ {
		combi[m] = l
		findElementsWithSum(arr, combi, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}

	return num
}

func main() {

	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}

	var addedSum int = 18

	var combinations [19]int

	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)

	//fmt.Println(check)

	//var check2 bool = findElement(arr,9)

	//fmt.Println(check2)

}

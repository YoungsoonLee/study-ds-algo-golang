package main

import "fmt"

// brute-force
func pairSum(arr []int, k int) {
	result := map[int]int{}
	for i := 0; i < len(arr); i++ {
		//currentEle := arr[i]
		target := k - arr[i]
		//fmt.Println(target)
		//fmt.Println(isExists(arr, 4))
		_, ok := result[target]
		if ok == false {
			fmt.Println(arr[i], target)
			result[arr[i]] = target
		}
	}

	fmt.Println(result)
}

func main() {
	a := []int{1, 3, 2, 2}
	k := 4

	pairSum(a, k)
}

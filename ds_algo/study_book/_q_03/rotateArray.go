package main

import "fmt"

func RotateArray(data []int, k int) {
	n := len(data)
	reverseArray(data, 0, k-1)
	reverseArray(data, k, n-1)
	reverseArray(data, 0, n-1)

}

func reverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	fmt.Println(data)
}

func myfunc(data []int, k int) {
	data.push(data[0 : k-1])
	fmt.Println(data)
}

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{3, 4, 5, 2, 6, 7, 8, 1, 9, 10} // ! not slice
	RotateArray(b, 5)

	myfunc(b, 5)
}

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintSlice(s)

	a := make([]int, 10) //slice
	PrintSlice(a)

	b := make([]int, 0, 10)
	PrintSlice(b) // [] :: len=0 cap=10

	c := s[0:4]
	PrintSlice(c) // [1 2 3 4] :: len=4 cap=10

	//z := []int{1, 2, 3}
	d := c[2:5]
	PrintSlice(d) // [3 4 5] :: len=3 cap=8 ?? 왜 5가 나올까?? 레인지에 값이 없으면 원래 포인터를 참고 하나. <----

	fmt.Printf("address of slice s %p \n", &s)
	fmt.Printf("address of slice c %p \n", &c)
	fmt.Printf("address of slice d %p \n", &d)
}

func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

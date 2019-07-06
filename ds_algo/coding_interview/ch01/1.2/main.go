package main

import "fmt"

func reverString(s string) string {
	size := len(s)
	// fmt.Println(size)
	// fmt.Println(string(s[size-1]))

	var rs string

	for i := size - 1; i >= 0; i-- {
		//fmt.Println(string(s[i]))
		rs = rs + string(s[i])
	}

	fmt.Println(rs)

	return rs
}

func main() {
	s := "hello world"
	fmt.Println(reverString(s))
}

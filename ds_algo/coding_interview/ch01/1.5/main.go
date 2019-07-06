package main

import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	var count int = 1
	var rs string
	for i := 1; i < len(s); i++ {
		//fmt.Println(string(s[i]), string(s[i+1]))
		if s[i-1] == s[i] {
			count++
		} else {
			rs = rs + string(s[i-1]) + strconv.Itoa(count)
			count = 1
		}

		//fmt.Println(i)
		//fmt.Println(len(s))
		if i+1 == len(s) { // check last
			//fmt.Println("???")
			rs = rs + string(s[i]) + strconv.Itoa(count)
		}
	}
	fmt.Println(rs)
	return rs
}

func main() {
	s := "aabccccccccaaa"
	compressString(s)
}

package main

import (
	"fmt"
)

func uniChar(s string) bool {
	r := make(map[string]int)

	for _, c := range s {
		// fmt.Println(string(c))
		if ok := r[string(c)]; ok == 1 {
			return false
		} else {
			r[string(c)] = 1
		}
	}

	return true
}

func main() {

	a := "aabbcc"
	b := "abcde"

	fmt.Println(uniChar(a))
	fmt.Println(uniChar(b))
}

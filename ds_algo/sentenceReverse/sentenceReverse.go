package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	words := strings.Fields(s)
	fmt.Println(words)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "     hello John    how are you     "
	//fmt.Println(strings.TrimSpace(s))
	fmt.Println(reverse(s))
}

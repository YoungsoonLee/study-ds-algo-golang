package main

import (
	"fmt"
	"strings"
)

func changeBlank(s string, size int) string {
	// change all space to '%20'
	var rs []string
	for i := 0; i < size; i++ {
		if string(s[i]) == " " {
			rs = append(rs, "%20")
		} else {
			rs = append(rs, string(s[i]))
		}
	}
	return strings.Join(rs, "")
}

func main() {
	s := "Mr Jone Smith     "
	fmt.Println(changeBlank(s, 13))
}

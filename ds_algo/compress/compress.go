package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func compress(s string) string {
	cnt := 1

	var buffer bytes.Buffer

	// fmt.Println(strconv.FormatFloat(3.14, 'f', 6, 64))

	var i int
	for i = 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt = cnt + 1
		} else {

			buffer.WriteString(string(s[i-1]))
			buffer.WriteString(strconv.Itoa(cnt))

			cnt = 1
		}
	}

	buffer.WriteString(string(s[i-1]))
	buffer.WriteString(strconv.Itoa(cnt))
	//fmt.Println("buffer: ", buffer.String())

	return buffer.String()
}

func main() {
	a := "AAAAABBBBccccaaa"
	fmt.Println(compress(a))
}

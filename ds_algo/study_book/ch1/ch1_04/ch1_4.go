package main

import "fmt"

func main() {
	s := "hello world"
	//s[0] = "H"
	r := []rune(s)
	fmt.Println(r)
	r[0] = 'H'
	s2 := string(r)
	fmt.Println(s2)
}

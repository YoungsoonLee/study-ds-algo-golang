package main

import (
	"fmt"
	"strconv"
)

type InterA interface {
	AAA(int) int
	BBB(int) string
}
type stA struct {
}

func (a *stA) AAA(x int) int {
	return x * x
}

func (a *stA) BBB(x int) string {
	return "x=" + strconv.Itoa(x)
}

func main() {
	var c InterA
	c = &stA{}

	fmt.Println(c.BBB(3))
}

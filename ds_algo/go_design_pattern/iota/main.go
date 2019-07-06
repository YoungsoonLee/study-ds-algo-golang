package main

import (
	"fmt"
)

func main() {
	type ByteSize float64

	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
	)

	fmt.Println(KB)
	//fmt.Printf("%b", int64(KB))
	//fmt.Println(strconv.FormatInt(int64(KB), 2))
}

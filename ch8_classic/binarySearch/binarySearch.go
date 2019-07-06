package main

import (
	"fmt"
	"sort"
)

func main() {
	els := []int{1, 3, 16, 10, 45, 31, 28, 36, 45, 75}
	target := 36

	i := sort.Search(len(els), func(i int) bool {
		return els[i] >= target
	})

	fmt.Println(i)

}

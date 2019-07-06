package main

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/set"
)

func findPair2(data []int, value int) bool {
	size := len(data)
	first := 0
	second := size - 1
	ret := false

	sort.Ints(data)

	for first < second {
		curr := data[first] + data[second]
		if curr == value {
			fmt.Println("the pair is ", data[first], data[second])
			ret = true
		} else if curr < value {
			first++
		} else {
			second--
		}
	}

	return ret
}

func findPair3(data []int, value int) bool {
	s := set.New()
	size := len(data)
	ret := false

	for i := 0; i < size; i++ {
		if s.Has(value - data[i]) {
			fmt.Println(i, "The pair is:", data[i], ",", (value - data[i]))
			ret = true
		} else {
			s.Insert(data[i])
		}
	}

	return ret
}

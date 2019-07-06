package main

import (
	"fmt"
	"sort"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 프린팅에 필요한 메서드 - 프린트하기 전에 요소들을 정렬함.
func (s Sequence) String() string {
	/*
		sort.Sort(s)
		str := "["
		for i, elem := range s {
			if i > 0 {
				str += " "
			}
			str += fmt.Sprint(elem)
			//str += string(elem)
		}
		return str + "]"
	*/
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}

func main() {
	a := &Sequence{1, 5, 2, 3, 7, 6, 8}
	fmt.Println(a.String())
}

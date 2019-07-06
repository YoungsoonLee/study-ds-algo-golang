package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	// fmt.Println(s)
	return strings.Join(s, "")
}

func anagram(s1, s2 string) bool {
	s1 = strings.Replace(s1, " ", "", -1)
	s2 = strings.Replace(s2, " ", "", -1)
	return sortString(s1) == sortString(s2)
}

func main() {
	fmt.Println(anagram("dog", "god"))
	fmt.Println(anagram("client eastwood", "olsd west action"))
	fmt.Println(anagram("aa", "bb"))
}

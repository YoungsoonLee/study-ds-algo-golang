// upper, lower ?
// blank?
// ascii
package main

import (
	"fmt"
	"sort"
	"strings"
)

func isPermutation(s string, t string) bool {
	// sort string
	// compare length
	if len(s) != len(t) {
		return false
	}

	s = strings.Replace(s, " ", "", -1) // 공백제거
	t = strings.Replace(t, " ", "", -1)
	return sortString(s) == sortString(t)
}

// !!!
func sortString(w string) string {
	s := strings.Split(w, "") // make array
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	//s := "dog"
	fmt.Println(isPermutation("dog", "god"))
}

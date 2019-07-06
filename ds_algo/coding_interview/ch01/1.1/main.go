// ascii = 256, unicode =
// blank ?
// a-zA-Z ?
// uppper lower ?
package main

import (
	"errors"
	"fmt"
)

//type input string // custom type
/*
func (i input) checkUniqueString() bool {
	fmt.Println("not yet")
	return true
}
*/

func isUniqueCharInString(s string) (bool, error) {
	tempMap := make(map[string]int)
	switch len(s) {
	case 0:
		return false, errors.New("empty string")
	case 1:
		return true, nil
	default:
		for _, c := range s {
			//fmt.Println(string(c))
			tempMap[string(c)]++
			if tempMap[string(c)] >= 2 {
				return false, nil
			}
		}
		/*
			for k, v := range tempMap {
				fmt.Printf("tempMap[%s] = %d\n", k, v)
				if v >= 2 {
					return false, nil
				}
			}
		*/
	}
	return true, nil
}

// solution1
func isUniqueCharInString2(s string) bool {
	if len(s) > 256 {
		return false
	}

	charSet := make([]bool, 256) // for ascii
	//fmt.Println(charSet)

	for _, c := range s {
		if charSet[c] {
			return false
		}
		charSet[c] = true
	}

	return true
}

//sort??

func main() {
	s := "Hello world"
	b, _ := isUniqueCharInString(s)
	fmt.Println(b)
	b = isUniqueCharInString2(s)
	fmt.Println(b)
}

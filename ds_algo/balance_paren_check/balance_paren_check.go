package main

import "fmt"

func isInArray(arr []string, item string) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}

	return false
}

func isMatch(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}

	return false
}

func balanceCheck(s string) bool {
	opening := []string{"[", "{", "("}
	matches := []string{"[]", "{}", "()"}

	stack := []string{}
	//need # edge process

	for _, p := range s {
		if isInArray(opening, string(p)) {
			stack = append(stack, string(p))
		} else {
			lastOpen := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			result := isMatch(matches, lastOpen+string(p))
			if result == false {
				return result
			}

		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(balanceCheck("{}([)"))
}

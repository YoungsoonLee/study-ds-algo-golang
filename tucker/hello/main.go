package main

import "fmt"

func Calculate(op string, a, b int) int {
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	}
	return 0

}

func main() {
	Test()
}

func Test() {
	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	}

	if !testCalculate("Test2", "+", 5, 4, 9) {
		return
	}

	if !testCalculate("Test3", "-", 5, 3, 2) {
		return
	}

	fmt.Println("Success")
}

func testCalculate(testcase, op string, a, b int, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed expexted: %d, output: %d\n", testcase, expected, o)
		return false
	}

	return true
}

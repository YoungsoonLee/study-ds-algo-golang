package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for index, val := range numbers {
		sum += val
		fmt.Println("[", index, ",", val, "]")
	}
	fmt.Println("\nSum: ", sum)

	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Println(k, "->", v)
	}

	fmt.Println(kvs)

	str := "Hello, World"
	for index, c := range str {
		fmt.Println(index, c)
	}

}

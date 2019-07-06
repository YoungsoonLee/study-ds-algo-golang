package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["apple"] = 40
	m["banana"] = 30
	m["mango"] = 50

	for key, val := range m {
		fmt.Print("[", key, "->", val, "]")
	}

	fmt.Println("apple price: ", m["apple"])
	delete(m, "apple")

	value, ok := m["apple"]
	fmt.Println("Apple price: ", value, "present: ", ok)

	value2, ok2 := m["banana"]
	fmt.Println("banana price: ", value2, "present: ", ok2)

}

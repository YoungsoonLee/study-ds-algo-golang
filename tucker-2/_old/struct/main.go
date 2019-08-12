package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := append(a, 3)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%p, %p", a, b)

	b[0] = 4
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	//a[2] = 3 // out of range
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%p, %p", a, b)

	var c = make([]int, len(a))
	copy(c, a)

	fmt.Println(c)

	fmt.Println(a)
	fmt.Println(c)

}

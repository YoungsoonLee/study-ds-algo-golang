package main

import "fmt"

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited")  // not
	fmt.Println("Exiting a()") // not
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()")
	fmt.Println("Exiting b()") // not
}

func main() {
	a()
	fmt.Println("main() ended!")
}

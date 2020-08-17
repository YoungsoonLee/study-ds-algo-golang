package main

import "fmt"

func pop(c chan int) {
	fmt.Println("pop")
	v := <-c

	fmt.Println(v)
}

func main() {
	var c chan int
	//c = make(chan int, 1)
	c = make(chan int) // zero channel

	go pop(c)
	c <- 10

	fmt.Println("end")
}

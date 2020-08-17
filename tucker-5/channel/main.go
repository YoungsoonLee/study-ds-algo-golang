package main

import "fmt"

func pop(c chan int) {
	fmt.Println("Pop func")
	v := <-c //waiting

	fmt.Println(v)

}

func main() {
	var c chan int
	c = make(chan int)

	//c <- 10 // blocking

	go pop(c)
	c <- 10

	fmt.Println("End program")
}

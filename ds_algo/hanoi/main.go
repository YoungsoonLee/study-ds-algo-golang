package main

import "fmt"

func TOHUtil(diskNum int, from string, destination string, buffer string) {
	if diskNum < 1 {
		return
	}

	TOHUtil(diskNum-1, from, buffer, destination)
	fmt.Println("Move disk ", diskNum, " from peg ", from, " to peg ", destination)
	TOHUtil(diskNum-1, buffer, destination, from)
}

func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	TOHUtil(num, "A", "C", "B")
}

func main() {
	TowersOfHanoi(5)
}

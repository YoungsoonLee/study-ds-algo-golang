package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)

	go push(c)

	t := time.After(10 * time.Second)
	t1 := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-t:
			fmt.Println("timeout")
			return
		case <-t1:
			fmt.Println("tick")

			/*
				default:
					fmt.Println("Idle")
					time.Sleep(1 * time.Second)
			*/
		}
	}

}

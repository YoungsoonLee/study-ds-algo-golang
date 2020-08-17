package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch Time: ", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference: ", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	londonTime := t.In(loc)
	fmt.Println("Paris: ", londonTime)

	for i := 0; i < 100; i++ {
		fmt.Println("Epoch Time: ", time.Now().Unix())
		time.Sleep(time.Second)
	}
}

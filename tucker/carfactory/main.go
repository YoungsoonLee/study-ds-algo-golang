package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChane chan Plane, outCarChan chan Car, ouPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += " tire c , "
			outCarChan <- car
		case plae := <-planeChane:
			plae.val += " tire p , "
			ouPlaneChan <- plae
		}
	}

}

func MakeEngine(carChan chan Car, planeChane chan Plane, outCarChan chan Car, ouPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += " Engine c , "
			outCarChan <- car
		case plae := <-planeChane:
			plae.val += " Engine p , "
			ouPlaneChan <- plae
		}
	}

}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car " + strconv.Itoa(i)}
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane " + strconv.Itoa(i)}
		i++
	}
}

func main() {
	carchan1 := make(chan Car)
	carchan2 := make(chan Car)
	carchan3 := make(chan Car)

	planechan1 := make(chan Plane)
	planechan2 := make(chan Plane)
	planechan3 := make(chan Plane)

	go StartCarWork(carchan1)
	go StartPlaneWork(planechan1)

	go MakeTire(carchan1, planechan1, carchan2, planechan2)
	go MakeEngine(carchan2, planechan2, carchan3, planechan3)

	//chan1 <- Car{val: "Car1: "}

	for {
		select {
		case result := <-carchan3:
			fmt.Println(result.val)
		case result := <-planechan3:
			fmt.Println(result.val)
		}

	}

}
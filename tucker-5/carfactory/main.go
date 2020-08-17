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

func MakeTire(carChan chan Car, planeChan chan Plane, outCar chan Car, outPlane chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outCar <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlane <- plane
		}
		/*
			car := <-carChan
			car.val += "Tire, "
			outCar <- car
		*/
	}

}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCar chan Car, outPlane chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCar <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlane <- plane
		}
		/*
			car := <-carChan
			car.val += "Engine, "
			outCar <- car
		*/
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
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartCarWork(carChan1)
	go StartPlaneWork(planeChan1)

	go MakeTire(carChan1, planeChan1, carChan2, planeChan2)
	go MakeEngine(carChan2, planeChan2, carChan3, planeChan3)

	for {
		select {
		case resultCar := <-carChan3:
			fmt.Println(resultCar.val)
		case resultPlane := <-planeChan3:
			fmt.Println(resultPlane.val)
		}
		/*
			resultCar := <-carChan3
			fmt.Println(resultCar.val)
		*/
	}

}

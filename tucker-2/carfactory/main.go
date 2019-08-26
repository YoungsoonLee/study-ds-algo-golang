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

func MakeTire(carChan chan Car, planeChan chan Plane, outChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlaneChan <- plane
		}
		/*
			car := <-carChan
			car.val += "Tire, "

			outChan <- car
		*/
	}

}

func MakeEngine(carChan chan Car, planeChan chan Plane, outChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}
		/*
			car := <-carChan
			car.val += "Engine, "

			outChan <- car
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
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartCarWork(chan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(chan1, planeChan1, chan2, planeChan2)
	go MakeEngine(chan2, planeChan2, chan3, planeChan3)

	for {
		select {
		case result := <-chan3:
			fmt.Println(result.val)
		case result := <-planeChan3:
			fmt.Println(result.val)
		}
		/*
			result := <-chan3
			fmt.Println(result.val)
		*/
	}

}

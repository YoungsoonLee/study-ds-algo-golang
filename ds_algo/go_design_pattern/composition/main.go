package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming")
}

func main() {
	swimmer := CompositeSwimmerA{MySwim: Swim}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	fish := Shark{Swim: Swim}
	fish.Eat()
	fish.Swim()

	swimmer2 := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer2.Train()
	swimmer2.Swim()
}

type Animal struct{}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal // embedding
	Swim   func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

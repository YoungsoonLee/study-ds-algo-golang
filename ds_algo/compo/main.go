package main

import "fmt"

type fuel int

const (
	GASOLINE fuel = iota
	BIO
	ELECTRIC
	JET
)

type vehicle struct {
	make  string
	model string
}

type engine struct {
	fuel   fuel
	thrust int
}

func (e *engine) start() {
	fmt.Println("Engine started.")
}

type truct struct {
	vehicle
	engine
	axels  int
	wheels int
	class  int
}

func (t *truct) druve() {
	fmt.Printf("Truck %s %s, on the go!\n", t.make, t.model)
}

type plane struct {
	vehicle
	engine
	engineCount int
	fixedWings  bool
	maxAltitude int
}

func (p *plane) fly() {
	fmt.Printf("aircraft %s %s clear for takeoff!\n", p.make, p.model)
}

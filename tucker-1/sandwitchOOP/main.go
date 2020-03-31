package main

import "fmt"

type Jam interface {
	GetOneSpoon() SpoonJam
}

type SpoonJam interface {
	String() string
}

type Bread struct {
	val string
}

type StJam struct {
}

type OrJam struct {
}

func (o *OrJam) GetOneSpoon() SpoonJam {
	return &SpoonOrJam{}

}

type SpoonStJam struct {
}

type SpoonOrJam struct {
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func (j *StJam) GetOneSpoon() *SpoonStJam {
	return &SpoonStJam{}
}

func (s *SpoonStJam) String() string {
	return "+ strawberry"
}

func (s *SpoonOrJam) String() string {
	return "+ Orange"
}

func main() {
	b := &Bread{}
	//j := &StJam{}
	j := &OrJam{}
	b.PutJam(j)
	fmt.Println(b.String())
}

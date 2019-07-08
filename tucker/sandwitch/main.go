package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}

	return breads
}

func OpendOpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func GetOneSpoon(jam *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func MakeeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

func main() {
	// 1. 빵 두개를 꺼낸다
	breads := GetBreads(2)

	// 2. 딸기쨈 뚜껑을 연다
	jam := &StrawberryJam{}
	OpendOpenStrawberryJam(jam)

	// 3. 딸기쨈을 한수푼 뜬다
	spoon := GetOneSpoon(jam)

	// 4. 딸기쨈을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}

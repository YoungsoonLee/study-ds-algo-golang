package main

import (
	"container/list" // 양방향 가진다.
	"fmt"
)

func main() {
	var intList list.List // !!!
	intList.PushBack(11)
	intList.PushBack(22)
	intList.PushBack(34)

	for el := intList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value.(int))
	}

}

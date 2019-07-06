package main

import "fmt"

type node struct {
	data int
	next *node
}

type List struct {
	head *node
}

func (l *List) add(data int) {
	if l.head == nil {
		l.head = &node{data: data, next: nil}
	} else {
		list := l.head

		for list.next != nil {
			list = list.next
		}
		list.next = &node{data: data, next: nil}
	}
}

func (l *List) remove(data int) {
	list := l.head
	if list.data == data { // first
		l.head = list.next
	} else {
		//list = list.next
		for list.next != nil {
			if list.next.data == data {
				//chect last
				if list.next.next == nil {
					list.next = nil
					return
				} else {
					list.next = list.next.next
				}
			}
			list = list.next
		}
	}
}

func (l *List) display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.data)
		list = list.next
	}

	fmt.Println()
}

func main() {
	l := List{}
	//l.display()
	l.add(5)
	//l.display()
	l.add(9)
	//l.display()
	l.add(13)
	//l.display()
	l.add(14)
	//l.display()
	l.add(15)
	l.display()
	l.remove(15)
	l.display()
}

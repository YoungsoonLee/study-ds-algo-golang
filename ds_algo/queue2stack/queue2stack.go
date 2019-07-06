package main

import (
	"fmt"
	"sync"
)

type queue2stack struct {
	lock     sync.Mutex
	inStack  []int
	outStack []int
}

func newQueue2Stack() *queue2stack {
	return &queue2stack{sync.Mutex{}, make([]int, 0), make([]int, 0)}
}

func (q *queue2stack) enqueue(e int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.inStack = append(q.inStack, e)
}

func (q *queue2stack) dequeue() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.outStack) == 0 {
		for _, v := range q.inStack {
			fmt.Println(v)
			q.outStack = append(q.outStack, v)
		}
	}

	r := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]

	return r
}

func main() {
	q := newQueue2Stack()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}

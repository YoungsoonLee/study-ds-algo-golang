package main

import (
	"fmt"
	"sync"
)

type queue struct {
	lock sync.Mutex
	q    []int
}

func newQueue() *queue {
	return &queue{sync.Mutex{}, make([]int, 0)}
}

func (q *queue) enqueue(v int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.q = append(q.q, v)
}

func (q *queue) dequeue() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	result := q.q[0]
	//fmt.Println(q.q)
	q.q = q.q[1:]
	//fmt.Println(q.q)
	return result
}

func (q *queue) isEmpty() bool {
	return len(q.q) == 0
}

func main() {
	q := newQueue()
	q.enqueue(1)
	q.enqueue(2)

	fmt.Println(q.dequeue())
	fmt.Println(q.isEmpty())
	fmt.Println(q.dequeue())
	fmt.Println(q.isEmpty())
}

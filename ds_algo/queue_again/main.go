package main

import (
	"errors"
	"fmt"
	"sync"
)

type queue struct {
	lock sync.Mutex
	data []int
}

func newQueue() *queue {
	return &queue{sync.Mutex{}, make([]int, 0)}
}

func (q *queue) enqueue(item int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data = append(q.data, item)
}

func (q *queue) dequeue() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.data) == 0 {
		return 0, errors.New("queue is empty") // return???
	}

	res := q.data[0]
	q.data = q.data[1:]
	return res, nil
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
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

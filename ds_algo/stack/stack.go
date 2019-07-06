package main

import (
	"errors"
	"fmt"
	"sync"
)

type stack struct {
	lock sync.Mutex
	s    []int
}

func newstack() *stack {
	return &stack{sync.Mutex{}, make([]int, 0)}
}

func (s *stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1] // remove at lice
	return res, nil
}

func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}

func (s *stack) Peek() int {
	return s.s[len(s.s)-1]
}

func main() {
	s := newstack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.isEmpty())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.isEmpty())
}

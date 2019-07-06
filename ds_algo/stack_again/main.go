package main

import (
	"errors"
	"fmt"
	"sync"
)

type stack struct {
	lock sync.Mutex
	data []int
}

func newStack() *stack {
	return &stack{sync.Mutex{}, make([]int, 0)}
}

func (s *stack) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// condition: data is positive int over 0
func (s *stack) Push(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, item)
}

func (s *stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	// check empty !!!
	l := len(s.data)
	if l == 0 {
		return 0, errors.New("empty stack")
	}

	res := s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) Peek() int {
	return s.data[len(s.data)-1]
}

func main() {
	s := newStack()
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

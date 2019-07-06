// 친밀도
package main

import (
	"fmt"
	"sync"

	"github.com/golang-collections/collections/set"
)

type queue struct {
	lock sync.Mutex
	q    []string
}

func newQueue() *queue {
	return &queue{sync.Mutex{}, make([]string, 0)}
}

func (q *queue) enqueue(v string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.q = append(q.q, v)
}

func (q *queue) dequeue() string {
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

type friend map[string][]string

func newFriend() *friend {
	f := &friend{
		"Summer": []string{"John", "Justin", "Mike"},
		"John":   []string{"Summer", "Justin"},
		"Justin": []string{"Summer", "John", "Mike", "May"},
		"Mike":   []string{"Summer", "Justin"},
		"May":    []string{"Justin", "Kim"},
		"Kim":    []string{"May"},
		"Tom":    []string{"Jerry"},
		"Jerry":  []string{"Tom"},
	}
	return f
}

func (f *friend) findFriend(item string) []string {
	//fmt.Println((*f)[item])
	return (*f)[item]
}

func (f *friend) findAllFriend(item string) {
	done := set.New()
	q := newQueue()

	done.Insert(item)
	q.enqueue(item)

	for !q.isEmpty() {
		s := q.dequeue()
		sf := f.findFriend(s)
		//fmt.Println(sf)

		for _, v := range sf {
			if !done.Has(v) {
				q.enqueue(v)
				done.Insert(v)
			}
		}

	}

	fmt.Println(q.q)
	fmt.Println(done)
}

func main() {
	f := newFriend()
	f.findAllFriend("Summer")
	f.findAllFriend("Jerry")
}

package queues

import (
	"fmt"
	"sync"
)

type Queue interface {
	Enqueue(item string) // add to queue
	Dequeue() string     // get from queue with deleting
	Front() string
	IsEmpty() bool
	Size() int
	PrintQueue()
}

type queue struct {
	items []string
	lock  sync.RWMutex
}

func (q *queue) PrintQueue() {
	fmt.Println(q.items)
}

func (q *queue) Enqueue(item string) {
	q.lock.Lock()
	q.items = append(q.items, item)
	q.lock.Unlock()
}

func (q *queue) Dequeue() string {
	var deletedElement string
	if q.IsEmpty() != true {
		q.lock.Lock()
		deletedElement = q.items[0]
		q.items = append(q.items[1:])
		q.lock.Unlock()
	}
	return deletedElement
}

func (q *queue) Front() string {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.items[0]
}

func (q *queue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.items) == 0
}

func (q *queue) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.items)
}

func NewQueue() Queue {
	queueObject := queue{items: []string{}, lock: sync.RWMutex{}}
	return &queueObject
}

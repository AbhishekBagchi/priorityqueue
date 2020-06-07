package priorityqueue

import (
	"container/heap"
	"errors"
)

//An Element is what's held by the priority queue
type Element struct {
	value    interface{}
	priority int
}

type elements []*Element

type PriorityQueue struct {
	queue pq_internal
}

func New(typ PriorityQueueType) *PriorityQueue {
	pq := &PriorityQueue{queue: pq_internal{typ: typ}}
	heap.Init(&pq.queue)
	return pq
}

type pq_internal struct {
	queue elements
	typ   PriorityQueueType
}

func (pq pq_internal) Len() int { return len(pq.queue) }

func MinQueueFunc(pi, pj int) bool {
	return pi < pj
}

func MaxQueueFunc(pi, pj int) bool {
	return pi > pj
}

func (pq pq_internal) Less(i, j int) bool {
	if pq.typ == MinQueue {
		return MinQueueFunc(pq.queue[i].priority, pq.queue[j].priority)
	} else {
		return MaxQueueFunc(pq.queue[i].priority, pq.queue[j].priority)
	}
}

func (pq pq_internal) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *pq_internal) Push(x interface{}) {
	element := x.(*Element)
	pq.queue = append(pq.queue, element)
}

func (pq *pq_internal) Top() Element {
	element := pq.queue[0]
	return *element
}

func (pq *pq_internal) Pop() interface{} {
	size := len(pq.queue)
	element := pq.queue[size-1]
	pq.queue[size-1] = nil //Avoid memory leak
	pq.queue = pq.queue[0 : size-1]
	return element
}

//Len returns the number of elements in the queue
func (pq *PriorityQueue) Len() int { return pq.queue.Len() }

func (pq *PriorityQueue) Push(value interface{}, priority int) {
	element := &Element{value, priority}
	heap.Push(&pq.queue, element)
}

func (pq *PriorityQueue) Pop() (interface{}, int, error) {
	if pq.Len() == 0 {
		return nil, 0, errors.New("No elements to pop")
	}
	element := *(heap.Pop(&pq.queue).(*Element))
	return element.value, element.priority, nil
}

func (pq *PriorityQueue) Top() (interface{}, int, error) {
	if pq.Len() == 0 {
		return nil, 0, errors.New("No elements to pop")
	}
	element := pq.queue.Top()
	return element.value, element.priority, nil
}

package priorityqueue

import (
	_ "container/heap"
)

//An Element is what's held by the priority queue
type Element struct {
	value    *interface{}
	priority int
}

type PriorityQueue struct {
	queue []*Element
	typ   PriorityQueueType
}

///
///type Interface interface {
///	// Len is the number of elements in the collection.
///	Len() int
///	// Less reports whether the element with
///	// index i should sort before the element with index j.
///	Less(i, j int) bool
///	// Swap swaps the elements with indexes i and j.
///	Swap(i, j int)
///}

func (pq PriorityQueue) Len() int { return len(pq.queue) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.typ == MinQueue {
		return pq.queue[i].priority < pq.queue[j].priority
	} else {
		return pq.queue[i].priority > pq.queue[j].priority
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

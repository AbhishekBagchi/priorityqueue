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

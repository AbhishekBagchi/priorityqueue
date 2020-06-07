package priorityqueue

type PriorityQueueType uint8

const (
	MinQueue PriorityQueueType = iota + 1 //Configures the queue to return the smallest element on pop
	MaxQueue PriorityQueueType = iota + 1 //Configures the queue to return the largest element on pop
)

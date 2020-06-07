package priorityqueue

import (
	"sort"
	"strconv"
	"testing"
)

func fillIntQueue(pq *PriorityQueue) int {
	priorities := map[int]int{
		1:  1,
		2:  10,
		3:  5,
		4:  3,
		5:  7,
		6:  6,
		7:  7,
		10: 9,
		21: 20,
		19: 22,
		17: 12,
	}
	for i, priority := range priorities {
		pq.Push(i, priority)
	}

	return len(priorities)
}

func TestPush(t *testing.T) {
	minQueue := New(MinQueue)
	maxQueue := New(MaxQueue)

	if minQueue.Len() != 0 {
		t.Error("Length of empty minQueue should be 0")
	}

	if maxQueue.Len() != 0 {
		t.Error("Length of empty maxQueue should be 0")
	}

	minSize := fillIntQueue(minQueue)
	maxSize := fillIntQueue(maxQueue)

	if minQueue.Len() != minSize {
		t.Error("Length of minQueue should be" + strconv.Itoa(minSize) + ", but it is " + strconv.Itoa(minQueue.Len()))
	}

	if maxQueue.Len() != maxSize {
		t.Error("Length of maxQueue should be" + strconv.Itoa(maxSize) + ", but it is " + strconv.Itoa(maxQueue.Len()))
	}
}

func TestPop(t *testing.T) {
	minQueue := New(MinQueue)
	maxQueue := New(MaxQueue)

	_, _, err := minQueue.Pop()
	if err == nil {
		t.Error("Pop from an empty queue should be an error")
	}

	_, _, err = maxQueue.Pop()
	if err == nil {
		t.Error("Pop from an empty queue should be an error")
	}

	fillIntQueue(minQueue)
	fillIntQueue(maxQueue)

	var maxResult []int
	for maxQueue.Len() > 0 {
		_, prt, err := maxQueue.Pop()
		if err != nil {
			t.Error("Pop should not have returned an error")
		}
		maxResult = append(maxResult, prt)
	}
	res := sort.SliceIsSorted(maxResult, func(i, j int) bool {
		return maxResult[i] > maxResult[j]
	})
	if res == false {
		t.Logf("%v", maxResult)
		t.Error("The elements in maxQueue should be sorted")
	}

	var minResult []int
	for minQueue.Len() > 0 {
		_, prt, _ := minQueue.Pop()
		minResult = append(minResult, prt)
	}
	res = sort.SliceIsSorted(minResult, func(i, j int) bool {
		return minResult[i] < minResult[j]
	})
	res = sort.SliceIsSorted(minResult, MinQueueFunc)
	if res == false {
		t.Logf("%v", minResult)
		t.Error("The elements in minQueue should be sorted")
	}
}

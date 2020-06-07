package priorityqueue

import (
	"sort"
	"strconv"
	"testing"
)

func TestPush(t *testing.T) {
	t.Parallel()
	minQueue := New(MinQueue)
	maxQueue := New(MaxQueue)
	if minQueue.Len() != 0 {
		t.Error("Length of empty minQueue should be 0")
	}

	if maxQueue.Len() != 0 {
		t.Error("Length of empty maxQueue should be 0")
	}

	priorities := [10]int{1, 10, 5, 3, 7, 9, 20, 22, 17, 12}
	for i, priority := range priorities {
		minQueue.Push(i, priority)
		maxQueue.Push(i, priority)
	}

	if minQueue.Len() != 10 {
		t.Error("Length of inQueue should be 10, but it is " + strconv.Itoa(minQueue.Len()))
	}

	var maxResult []int
	for maxQueue.Len() > 0 {
		_, prt, _ := maxQueue.Pop()
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

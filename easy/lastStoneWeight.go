package easy

import (
	"container/heap"
	"fmt"
)

type Intheap []int

func (h *Intheap) Push(x interface{}) {
	*h = append(*h,x.(int))
}

func (h *Intheap) Pop() interface{} {
	old := *h
	item := len(old) - 1
	x := old[item]
	*h = old[:item]
	return x
}

func (h *Intheap) Len() int {
	return len(*h)
}

func (h *Intheap) Less(i,j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *Intheap) Swap(i,j int) {
	(*h)[i],(*h)[j] = (*h)[j],(*h)[i]
}


type IntMaxHeap []int

func (pq *IntMaxHeap) Len() int {
	return len(*pq)
}
func (pq *IntMaxHeap) Less(i, j int) bool {
	return (*pq)[i] > (*pq)[j]
}
func (pq *IntMaxHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *IntMaxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *IntMaxHeap) Pop() interface{} {
	n := len(*pq) - 1
	x := (*pq)[n]
	*pq = (*pq)[:n]
	return x
}
func (pq *IntMaxHeap) Peek() int {
	return (*pq)[0]
}

func lastStoneWeight(stones []int) int {
	h := &Intheap{}
	for _,v := range stones {
		heap.Push(h,v)
	}
	fmt.Printf("%v\n",h)

	for h.Len() > 1 {
		big := heap.Pop(h).(int)
		small := heap.Pop(h).(int)
		fmt.Printf("big:%d,small:%d\n",big,small)
		if big > small {
			heap.Push(h,big-small)
		}
	}

	if h.Len() > 0 {
		fmt.Printf("%v\n",h)
		return heap.Pop(h).(int)
	}

	return 0
}

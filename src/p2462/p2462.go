package p2462

import (
	"container/heap"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h Heap) First() int {
	return h[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap := &Heap{}
	rightHeap := &Heap{}

	left := 0
	right := len(costs) - 1

	for i := 0; i < candidates; i++ {
		heap.Push(leftHeap, costs[i])
		left += 1
	}

	for i := len(costs) - 1; i >= max(left, len(costs)-candidates); i-- {
		heap.Push(rightHeap, costs[i])
		right--
	}

	total := 0

	for i := 0; i < k; i++ {
		if leftHeap.Len() > 0 && rightHeap.Len() > 0 && leftHeap.First() <= rightHeap.First() || rightHeap.Len() == 0 {
			total += heap.Pop(leftHeap).(int)
			if left <= right {
				heap.Push(leftHeap, costs[left])
				left++
			}
		} else {
			total += heap.Pop(rightHeap).(int)
			if left <= right {
				heap.Push(rightHeap, costs[right])
				right--
			}
		}
	}

	return int64(total)
}

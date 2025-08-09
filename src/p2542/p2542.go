package p2542

import (
	"container/heap"
	"sort"
)

type Pair struct {
	A int
	B int
}

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
	*h = old[:n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	pairs := make([]Pair, 0, len(nums1))

	for i := 0; i < len(nums1); i++ {
		pair := Pair{A: nums1[i], B: nums2[i]}
		pairs = append(pairs, pair)
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].B > pairs[j].B
	})

	h := &Heap{}
	total := 0

	for i := 0; i < k; i++ {
		heap.Push(h, pairs[i].A)
		total += pairs[i].A
	}

	result := total * pairs[k-1].B

	for i := k; i < len(pairs); i++ {
		total -= heap.Pop(h).(int)
		total += pairs[i].A
		heap.Push(h, pairs[i].A)
		result = max(result, total*pairs[i].B)
	}

	return int64(result)
}

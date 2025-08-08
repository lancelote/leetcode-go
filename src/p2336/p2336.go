package p2336

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

type SmallestInfiniteSet struct {
	Left      int
	Heap      *Heap
	AddedBack map[int]struct{}
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		Left:      1,
		Heap:      &Heap{},
		AddedBack: make(map[int]struct{}),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var r int

	if this.Heap.Len() > 0 && this.Heap.First() < this.Left {
		r = heap.Pop(this.Heap).(int)
		delete(this.AddedBack, r)
		return r
	}
	r = this.Left
	this.Left++
	return r
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.Left {
		if _, ok := this.AddedBack[num]; !ok {
			heap.Push(this.Heap, num)
			this.AddedBack[num] = struct{}{}
		}
	}
}

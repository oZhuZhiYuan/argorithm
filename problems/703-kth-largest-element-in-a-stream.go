/*
problem 703
https://leetcode.com/problems/kth-largest-element-in-a-stream/
*/
package problems

import "container/heap"

/*  implement int head */
// https://studygolang.com/pkgdoc
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/* solution down here */
type KthLargest struct {
	PQ      IntHeap
	MaxSize int
}

func KthConstructor(k int, nums []int) KthLargest {
	kth := KthLargest{}
	for _, v := range nums {
		kth.PQ = append(kth.PQ, v)
	}
	heap.Init(&kth.PQ)
	for kth.PQ.Len() > k {
		heap.Pop(&kth.PQ)
	}
	kth.MaxSize = k
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.PQ.Len() == 0 {
		heap.Push(&this.PQ, val)
		return this.PQ[0]
	}
	// when PQ is full, exec this
	if this.PQ.Len() == this.MaxSize {
		if val < this.PQ[0] {
			return this.PQ[0]
		}
		heap.Pop(&this.PQ)
	}
	// when PQ is not full, just push
	heap.Push(&this.PQ, val)
	return this.PQ[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

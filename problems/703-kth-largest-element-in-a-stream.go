/*
problem 703
https://leetcode.com/problems/kth-largest-element-in-a-stream/
问题：
返回数据流中第k大的数
解法：
维护一个mini heap(小顶堆)，每进来一个int，就与堆顶比较
如果小于堆顶，则返回堆顶的值
如果大于堆顶，则push进堆，然后返回堆顶的值
*/
package problems

import "container/heap"

/*
第一步 实现 heap.Interface
https://studygolang.com/pkgdoc
*/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/*
第二步 以下是解答
*/
type KthLargest struct {
	PQ IntHeap
	K  int
}

func KthConstructor(k int, nums []int) KthLargest {
	var kth KthLargest
	kth.K = k
	kth.PQ = nums
	heap.Init(&kth.PQ)
	for len(kth.PQ) > k {
		heap.Pop(&kth.PQ)
	}
	return kth
}

func (kth *KthLargest) Add(val int) int {
	if kth.PQ.Len() < kth.K {
		heap.Push(&kth.PQ, val)
		return kth.PQ[0]
	}
	if val > kth.PQ[0] {
		heap.Pop(&kth.PQ)
		heap.Push(&kth.PQ, val)
	}
	return kth.PQ[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

package testcase

import (
	pro "argo/problems"
	util "argo/utils"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	testcase := []int{1, 2, 3, 4, 5}
	expect := []int{2, 1, 4, 3, 5}
	head := createLinkedList(testcase)
	newHead := pro.SwapPairs(head)
	ret := []int{}
	for newHead != nil {
		ret = append(ret, newHead.Val)
		newHead = newHead.Next
	}
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test SwapPairs fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

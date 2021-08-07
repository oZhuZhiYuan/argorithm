package testcase

import (
	pro "argo/problems"
	util "argo/utils"
	"testing"
)

func TestReverseList(t *testing.T) {
	testcase := []int{1, 2, 3}
	expect := []int{3, 2, 1}
	head := createLinkedList(testcase)
	newHead := pro.ReverseList(head)
	ret := []int{}
	for newHead != nil {
		ret = append(ret, newHead.Val)
		newHead = newHead.Next
	}
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test ReverseList fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test ReverseList success\nexpect: %d\nresult: %d", expect, ret)
}

package testcase

import (
	arr "argo/argorithm"
	util "argo/utils"
	"testing"
)

// 给定一个数组，生成int 链表
func createLinkedList(a []int) *arr.ListNode {
	if len(a) == 0 {
		return nil
	}
	headnode := arr.ListNode{
		Val:  a[0],
		Next: nil,
	}
	cur := &headnode
	for i := 1; i < len(a); i++ {
		node := arr.ListNode{
			Val:  a[i],
			Next: nil,
		}
		cur.Next = &node
		cur = cur.Next
	}
	return &headnode
}

func TestReverseList(t *testing.T) {
	testcase := []int{1, 2, 3}
	expect := []int{3, 2, 1}
	head := createLinkedList(testcase)
	newHead := arr.ReverseList(head)
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

func TestSwapPairs(t *testing.T) {
	testcase := []int{1, 2, 3, 4, 5}
	expect := []int{2, 1, 4, 3, 5}
	head := createLinkedList(testcase)
	newHead := arr.SwapPairs(head)
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

func TestHasCycle(t *testing.T) {
	testcase := []int{3, 2, 0, -4}
	pos := 1
	expect := true
	head := createLinkedList(testcase)
	// make a cycle on a list
	p, cur, i := head, head, 0
	for cur.Next != nil {
		if i == pos {
			p = cur
		}
		cur = cur.Next
		i++
	}
	if pos != -1 {
		cur.Next = p
	}

	ret := arr.HasCycle(head)
	if ret != expect {
		t.Fatalf("test HasCycle fail\nexpect: %t\nresult: %t", expect, ret)
	}

	t.Logf("test HasCycle success\nexpect: %t\nresult: %t", expect, ret)
}

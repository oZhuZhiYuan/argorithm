package testcase

import (
	pro "argo/problems"
)

// 给定一个数组，生成int 链表
func createLinkedList(a []int) *pro.ListNode {
	if len(a) == 0 {
		return nil
	}
	headnode := pro.ListNode{
		Val:  a[0],
		Next: nil,
	}
	cur := &headnode
	for i := 1; i < len(a); i++ {
		node := pro.ListNode{
			Val:  a[i],
			Next: nil,
		}
		cur.Next = &node
		cur = cur.Next
	}
	return &headnode
}

package problems

// https://leetcode-cn.com/problems/reverse-linked-list/
func ReverseList(head *ListNode) *ListNode {
	var cur, pre *ListNode
	cur, pre = head, nil
	for cur != nil {
		// tmp := cur.Next
		// cur.Next = pre
		// pre = cur
		// cur = tmp
		cur.Next, cur, pre = pre, cur.Next, cur
	}
	return pre
}

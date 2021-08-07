package problems

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func SwapPairs(head *ListNode) *ListNode {
	newHead := ListNode{} // head 前一个节点
	pre := &newHead
	pre.Next = head
	for pre.Next != nil && pre.Next.Next != nil {
		// define 2 pointers to store last and next nodes
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
	}
	return newHead.Next
}

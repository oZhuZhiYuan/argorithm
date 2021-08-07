package argorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

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

// https://leetcode-cn.com/problems/linked-list-cycle/
// hash set
func HasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]int)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur], cur = 1, cur.Next
	}
	return false
}

// fast slow pointer
func HasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && slow != nil {
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

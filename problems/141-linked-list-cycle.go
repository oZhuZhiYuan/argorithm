package problems

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

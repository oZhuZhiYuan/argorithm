package testcase

import (
	pro "argo/problems"
	"testing"
)

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

	ret := pro.HasCycle(head)
	if ret != expect {
		t.Fatalf("test HasCycle fail\nexpect: %t\nresult: %t", expect, ret)
	}

	t.Logf("test HasCycle success\nexpect: %t\nresult: %t", expect, ret)
}

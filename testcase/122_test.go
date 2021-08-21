package testcase

import (
	"argo/problems"
	"testing"
)

func TestMaxProfit(t *testing.T) {

	testcase := []int{1}
	expect := 0
	ret := problems.MaxProfit(testcase)
	if expect != ret {
		t.Fatalf("test MaxProfit fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test MaxProfit success\nexpect: %d\nresult: %d", expect, ret)

	testcase = []int{7, 1, 5, 3, 6, 4}
	expect = 7
	ret = problems.MaxProfit(testcase)
	if expect != ret {
		t.Fatalf("test MaxProfit fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test MaxProfit success\nexpect: %d\nresult: %d", expect, ret)

	testcase = []int{1, 2, 3, 4, 5}
	expect = 4
	ret = problems.MaxProfit(testcase)
	if expect != ret {
		t.Fatalf("test MaxProfit fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test MaxProfit success\nexpect: %d\nresult: %d", expect, ret)

	testcase = []int{7, 6, 4, 3, 1}
	expect = 0
	ret = problems.MaxProfit(testcase)
	if expect != ret {
		t.Fatalf("test MaxProfit fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test MaxProfit success\nexpect: %d\nresult: %d", expect, ret)

}

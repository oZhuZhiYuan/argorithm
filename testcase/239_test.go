package testcase

import (
	pro "argo/problems"
	util "argo/utils"
	"testing"
)

// go test -v -run MaxSlidingWindow
func TestMaxSlidingWindow1(t *testing.T) {
	nums, k := []int{1, 3, -1, -3, 5, 3, 6, 7, 4, 3, 3}, 3
	expect := []int{3, 3, 5, 5, 6, 7, 7, 7, 4}
	ret := pro.MaxSlidingWindow(nums, k)
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test SwapPairs fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

func TestMaxSlidingWindow2(t *testing.T) {
	nums, k := []int{1}, 1
	expect := []int{1}
	ret := pro.MaxSlidingWindow(nums, k)
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test SwapPairs fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

func TestMaxSlidingWindow3(t *testing.T) {
	nums, k := []int{1, -1}, 1
	expect := []int{1, -1}
	ret := pro.MaxSlidingWindow(nums, k)
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test SwapPairs fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

func TestMaxSlidingWindow4(t *testing.T) {
	nums, k := []int{9, 11}, 2
	expect := []int{11}
	ret := pro.MaxSlidingWindow(nums, k)
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test SwapPairs fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

func TestMaxSlidingWindow5(t *testing.T) {
	nums, k := []int{4, -2}, 2
	expect := []int{4}
	ret := pro.MaxSlidingWindow(nums, k)
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test SwapPairs fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

func TestMaxSlidingWindow6(t *testing.T) {
	nums, k := []int{7, 2, 4}, 2
	expect := []int{7, 4}
	ret := pro.MaxSlidingWindow(nums, k)
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test SwapPairs fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

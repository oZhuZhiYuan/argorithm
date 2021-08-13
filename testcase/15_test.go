package testcase

import (
	pro "argo/problems"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expect := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	ret := pro.ThreeSum(nums)
	if len(expect) != len(ret) {
		t.Fatalf("test ThreeSum fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

func TestThreeSum1(t *testing.T) {
	nums := []int{}
	expect := []int{}
	ret := pro.ThreeSum(nums)
	if len(expect) != len(ret) {
		t.Fatalf("test ThreeSum fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

func TestThreeSum2(t *testing.T) {
	nums := []int{0}
	expect := []int{}
	ret := pro.ThreeSum(nums)
	if len(expect) != len(ret) {
		t.Fatalf("test ThreeSum fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

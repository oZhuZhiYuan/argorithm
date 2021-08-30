package testcase

import (
	"argo/problems"
	"testing"
)

func TestMySqrt(t *testing.T) {
	// case 1
	testcase := 4
	ret := problems.MySqrt(testcase)
	expect := 2
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
	// case 2
	testcase = 8
	ret = problems.MySqrt(testcase)
	expect = 2
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
	// case 3
	testcase = 24
	ret = problems.MySqrt(testcase)
	expect = 4
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)

	// case 4
	testcase = 2
	ret = problems.MySqrt(testcase)
	expect = 1
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %d\nresult: %d", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %d\nresult: %d", expect, ret)
}

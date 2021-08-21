package testcase

import (
	"argo/problems"
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	// case 1
	expect := math.Pow(2, 10)
	ret := problems.MyPow(2, 10)
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %f\nresult: %f", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %f\nresult: %f", expect, ret)
	// case 2
	expect = math.Pow(2.1, 3)
	ret = problems.MyPow(2.1, 3)
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %f\nresult: %f", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %f\nresult: %f", expect, ret)
	// case3
	expect = math.Pow(2, -2)
	ret = problems.MyPow(2, -2)
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %f\nresult: %f", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %f\nresult: %f", expect, ret)

	// case4
	expect = math.Pow(2, 51)
	ret = problems.MyPow(2, 51)
	if expect != ret {
		t.Fatalf("test ThreeSum fail\nexpect: %f\nresult: %f", expect, ret)
	}
	t.Logf("test SwapPairs success\nexpect: %f\nresult: %f", expect, ret)
}

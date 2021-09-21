package testcase

import (
	"argo/problems"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	// case 1
	var num uint32 = 6
	expect := 2
	ret := problems.HammingWeight(num)
	if ret != expect {
		t.Fatalf("test HammingWeight fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test HammingWeight success\nexpect: %d\nresult: %d", expect, ret)
	// case 2
	num = 1
	expect = 1
	ret = problems.HammingWeight(num)
	if ret != expect {
		t.Fatalf("test HammingWeight fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test HammingWeight success\nexpect: %d\nresult: %d", expect, ret)
	// case 3
	num = 10
	expect = 2
	ret = problems.HammingWeight(num)
	if ret != expect {
		t.Fatalf("test HammingWeight fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test HammingWeight success\nexpect: %d\nresult: %d", expect, ret)
}

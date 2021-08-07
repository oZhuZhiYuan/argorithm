package testcase

import (
	argo "argo/argorithm"
	"testing"
)

func TestBracketIsVaild(t *testing.T) {
	testcase := "{{}{}{}[]()[]}"
	expect := true
	ret := argo.BracketIsValid(testcase)
	if ret != expect {
		t.Fatalf("test BracketIsVaild fail\nexpect: %t\nresult: %t", expect, ret)
	}

	t.Logf("test BracketIsVaild success\nexpect: %t\nresult: %t", expect, ret)
}

package testcase

import (
	pro "argo/problems"
	"testing"
)

func TestBracketIsVaild(t *testing.T) {
	testcase := "{{}{}{}[]()[]}"
	expect := true
	ret := pro.BracketIsValid(testcase)
	if ret != expect {
		t.Fatalf("test BracketIsVaild fail\nexpect: %t\nresult: %t", expect, ret)
	}

	t.Logf("test BracketIsVaild success\nexpect: %t\nresult: %t", expect, ret)
}

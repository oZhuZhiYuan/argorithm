package testcase

import (
	pro "argo/problems"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	ss := "ab"
	tt := "abc"
	expect := false
	ret := pro.IsAnagram(ss, tt)
	if ret != expect {
		t.Fatalf("test IsAnagram fail\nexpect: %t\nresult: %t", expect, ret)
	}

	t.Logf("test IsAnagram success\nexpect: %t\nresult: %t", expect, ret)

}

func TestIsAnagram1(t *testing.T) {
	ss := "anagram"
	tt := "nagaram"
	expect := true
	ret := pro.IsAnagram(ss, tt)
	if ret != expect {
		t.Fatalf("test IsAnagram fail\nexpect: %t\nresult: %t", expect, ret)
	}

	t.Logf("test IsAnagram success\nexpect: %t\nresult: %t", expect, ret)

}

func TestIsAnagram2(t *testing.T) {
	ss := "rat"
	tt := "car"
	expect := false
	ret := pro.IsAnagram(ss, tt)
	if ret != expect {
		t.Fatalf("test IsAnagram fail\nexpect: %t\nresult: %t", expect, ret)
	}

	t.Logf("test IsAnagram success\nexpect: %t\nresult: %t", expect, ret)

}

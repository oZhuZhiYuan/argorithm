package testcase

import (
	"argo/problems"
	"testing"
)

func TestNumIslands(t *testing.T) {
	testcase := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	expect := 3
	ret := problems.NumIslands(testcase)
	if ret != expect {
		t.Fatalf("test NumIslands fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test NumIslands success\nexpect: %d\nresult: %d", expect, ret)
}

func TestNumIslands1(t *testing.T) {
	testcase := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	expect := 3
	ret := problems.NumIslands1(testcase)
	if ret != expect {
		t.Fatalf("test NumIslands1 fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test NumIslands1 success\nexpect: %d\nresult: %d", expect, ret)
}

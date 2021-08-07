package testcase

import (
	pro "argo/problems"
	"testing"
)

func TestTwosum(t *testing.T) {
	a := []int{2, 5, 7, 10, 11}
	target := 15
	ret := pro.Twosum(a, target)
	if a[ret[0]]+a[ret[1]] != target {
		t.Fatalf("Twosum test fail")
	}
	t.Logf("Twosum test success")
}

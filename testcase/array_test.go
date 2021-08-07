package testcase

import (
	arr "argo/argorithm"
	"testing"
)

//  go test -v -run TestFibonacci
func TestFibonacci(t *testing.T) {
	// var n uint64
	n := 40
	if arr.Fibonacci(uint64(n)) != 102334155 {
		t.Fatalf("Fibonacci(%d) test fail", n)
	}
	t.Logf("Fibonacci(%d) test success", n)
}

func TestFibonacci2(t *testing.T) {
	n := 50
	if arr.Fibonacci2(uint64(n)) != 12586269025 {
		t.Fatalf("Fibonacci2(%d) test fail", n)
	}
	t.Logf("Fibonacci2(%d) test success", n)
}

func TestTwosum(t *testing.T) {
	a := []int{2, 5, 7, 10, 11}
	target := 15
	ret := arr.Twosum(a, target)
	if a[ret[0]]+a[ret[1]] != target {
		t.Fatalf("Twosum test fail")
	}
	t.Logf("Twosum test success")
}

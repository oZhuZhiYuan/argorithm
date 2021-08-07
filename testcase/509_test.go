package testcase

import (
	pro "argo/problems"
	"testing"
)

//  go test -v -run TestFibonacci
func TestFibonacci(t *testing.T) {
	// var n uint64
	n := 40
	if pro.Fibonacci(uint64(n)) != 102334155 {
		t.Fatalf("Fibonacci(%d) test fail", n)
	}
	t.Logf("Fibonacci(%d) test success", n)
}

func TestFibonacci2(t *testing.T) {
	n := 50
	if pro.Fibonacci2(uint64(n)) != 12586269025 {
		t.Fatalf("Fibonacci2(%d) test fail", n)
	}
	t.Logf("Fibonacci2(%d) test success", n)
}

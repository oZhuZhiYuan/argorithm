package problems

import "math"

// fibonacci O(2^n)
func Fibonacci(n uint64) uint64 {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// fibonacci O(n)
func Fibonacci2(n uint64) uint64 {
	// 0 1 1 2 3 5 8
	if n == 0 || n == 1 {
		return n
	}
	var b2, b1, i, ret uint64 = 0, 1, 2, 0
	for i = 2; i <= n; i++ {
		ret = b1 + b2
		b2, b1 = b1, ret
	}
	return ret
}

// leetcode 0 <= n <= 30
func Fibonacci3(n int) int {
	// 0 1 1 2 3 5 8
	if n == 0 || n == 1 {
		return n
	}
	b1, b2, ret, i := 1, 0, 0, 2
	for i = 2; i <= n; i++ {
		ret = b1 + b2
		b2 = b1
		b1 = ret
	}
	return ret
}

// fibonacci matrix 通项公式解法，精度不够
func Fibonacci4(n int) int {
	sqrt5 := math.Sqrt(5)
	nn := float64(n)
	ret := 1 / sqrt5 * (math.Pow(((1+sqrt5)/2), nn) - math.Pow(((1-sqrt5)/2), nn))
	return int(ret)
}

// 动态规划法

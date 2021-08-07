package argorithm

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
	var b1, b2, ret, i uint64 = 1, 0, 0, 2
	for i = 2; i <= n; i++ {
		ret = b1 + b2
		b2 = b1
		b1 = ret
	}
	return ret
}

// fibonacci O(log(n))

// https://leetcode-cn.com/problems/two-sum/
func Twosum(a []int, target int) [2]int {
	m := make(map[int]int)
	for i := range a {
		x := target - a[i]
		if _, ok := m[x]; ok {
			return [2]int{m[x], i}
		}
		m[a[i]] = i
	}
	return [2]int{-1, -1}
}

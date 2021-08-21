package problems

import (
	"math"
)

func MyPow1(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func MyPow(x float64, n int) float64 {
	// recursion terminator
	switch n {
	case 0:
		return 1
	case 1:
		return x
	case 2:
		return x * x
	}
	// 分治条件和子问题递归
	if n < 0 {
		return 1 / MyPow(x, -n)
	}
	if n&1 == 0 {
		ret := MyPow(x, n>>1)
		// generate the final result
		return ret * ret
	}
	ret := MyPow(x, n>>1)
	return ret * ret * x
}

// best
func MyPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow float64 = 1
	for n != 0 {
		if n&1 != 0 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}

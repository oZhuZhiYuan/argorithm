package problems

/*
算法思想：分治
如果n是偶数，算一半，将一半的结果相乘
如果n是奇数，算(n-1)/2,将结果相乘再乘x

时间复杂度为O(logn)
*/
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
	//  n < 0 ,  return reciprocal
	if n < 0 {
		return 1 / MyPow(x, -n)
	}
	// n is even
	if n&1 == 0 {
		ret := MyPow(x, n>>1)
		return ret * ret
	}
	// n is odd
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

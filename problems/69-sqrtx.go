package problems

/*
https://leetcode-cn.com/problems/sqrtx/
解法1  binary search 二分查找，需要考虑防止溢出
解法2  牛顿迭代法
*/

func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	ret, left, right := 0, 0, x
	for left <= right {
		// mid = (left + right)/2 , 可能会溢出
		mid := left + (right-left)>>1
		// mid*mid 可能会溢出
		if mid == x/mid {
			return mid
		}
		if mid > x/mid {
			right = mid - 1
			continue
		}
		if mid < x/mid {
			ret = mid
			left = mid + 1
			continue
		}
	}
	return ret
}

func MySqrt1(x int) int {
	ret := x
	for ret*ret > x {
		ret = (ret + x/ret) >> 1
	}
	return ret
}

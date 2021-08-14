package problems

/*
解法：
每次取一个数a，从其后的数组中找到两个数只和为-a的，加入到结果中
[1 0 -1 1 0 1 0]
排序：
[-1 0 0 0 1 1 1]
第一次循环：
[(-1)[0 0 0 1 1 1]] ,find 2sum==1
第二次循环：
[-1 (0)[0 0 1 1 1]] ,find 2sum==0
...
*/

import (
	"sort"
)

func ThreeSum1(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	rem := make(map[[3]int]int)
	for i := 0; i < len(nums)-2; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		m := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			if _, ok := m[nums[j]]; !ok {
				m[-nums[i]-nums[j]] = 1
			} else {
				rem[[3]int{nums[i], nums[j], -nums[i] - nums[j]}] = 1
			}
		}
	}
	for k := range rem {
		v := make([]int, 3)
		copy(v, k[:])
		ret = append(ret, v)
	}
	return ret
}

// nums1 := []int{1, 0, -1, 1, 0, 1, 0}
func ThreeSum(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 3 {
		return ret
	}
	// 排序，使得重复的结果数组顺序相同，便于去重
	sort.Ints(nums)
	retm := make(map[[3]int]int)
	for i, a := range nums[:len(nums)-2] {
		// 如果a重复出现，就直接跳过。因为a已经从其后的数组中find一次2sum了。
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		m := make(map[int]int)
		// 从a后的数组中查找2sum = -a
		for _, b := range nums[i+1:] {
			if _, ok := m[-a-b]; ok {
				retm[[3]int{a, b, -a - b}] = 1
			}
			m[b] = 1
		}
	}
	for k := range retm {
		r := make([]int, 3)
		copy(r, k[:])
		ret = append(ret, r)
	}
	return ret
}

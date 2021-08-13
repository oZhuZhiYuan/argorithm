package problems

import "sort"

func ThreeSum(nums []int) [][]int {
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

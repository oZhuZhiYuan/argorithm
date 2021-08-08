package problems

import (
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

// 使用数组排序
func MaxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	var win IntSlice
	ret := []int{}
	// 先实现窗口移动
	for i := k; i <= len(nums); i++ {
		win = nums[i-k : i]
		sort.Sort(win)
		ret = append(ret, win[len(win)-1])
	}
	// 选出窗口中最大的值
	return ret
}

// 使用deque 双向pop队列
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	// win 记录元素下标，ret记录每次移动最大元素
	win, ret := []int{}, []int{}
	for i, v := range nums {
		// 窗口向右移动，pop最左边的成员
		if i >= k && win[0] <= i-k {
			win = win[1:]
		}
		// 新的元素v入队列前，将所有比新元素v小的成员从右边pop，新元素下标加入QUEUE
		for len(win) > 0 && nums[win[len(win)-1]] < v {
			win = win[:len(win)-1]
		}
		win = append(win, i)
		// 只有元素完全进入窗口时，才开始记录最大值
		if i >= k-1 {
			ret = append(ret, nums[win[0]])
		}
	}
	return ret
}

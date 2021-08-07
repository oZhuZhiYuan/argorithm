package problems

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

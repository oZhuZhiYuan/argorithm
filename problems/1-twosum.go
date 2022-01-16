package problems

/* https://leetcode-cn.com/problems/two-sum/
返回两数之和的【下标】
*/

// O(n)
func Twosum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{0, 0}
	}
	map1 := make(map[int]int)
	for i, v := range nums {
		x := target - v
		if j, ok := map1[x]; ok {
			return []int{j, i}
		}
		map1[v] = i
	}
	return []int{0, 0}
}

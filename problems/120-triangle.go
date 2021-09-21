package problems

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

解法：动态规划，（反向）递推
dp[i,j]: 以dp[i,j]为始，到底部的最小路径和。
公式: dp[i,j]=min(dp[i-1,j],dp[i-1,j+1]) + triangle[i][j]
简化:
1.优化空间复杂度，只需记录上一行的dp[i]
2.优化空间复杂度，直接将dp[i,j保存到triangle[i][j]


*/

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 1.优化空间复杂度，只需记录上一行的dp[i]
func MinimumTotal(triangle [][]int) int {
	if len(triangle) < 2 {
		return triangle[0][0]
	}
	// 最后一行的值，都是dp[i,j]
	mini := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := range triangle[i] {
			mini[j] = minInt(mini[j], mini[j+1]) + triangle[i][j]
		}
	}
	return mini[0]
}

// 2.优化空间复杂度，直接将dp[i,j保存到triangle[i][j]
func MinimumTotal1(triangle [][]int) int {
	if len(triangle) < 2 {
		return triangle[0][0]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triangle[i][j] += minInt(triangle[i+1][j], triangle[i+1][j+1])
		}

	}
	return triangle[0][0]
}

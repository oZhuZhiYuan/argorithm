package problems

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

链接：https://leetcode-cn.com/problems/climbing-stairs


解法：动态规划（反向）递推
状态定义:  opt[n] = 第n阶的走法总和
转移方程: opt(n) = opt(n-1) + opt(n-2)
*/

// 简单写法 O(n), 动态规划变种，只需记录前两个值，同时省略一个参数
func ClimbStairs(n int) int {
	// opt(n) = opt(n-1) + opt(n-2)
	b1, b2 := 1, 1
	for i := 2; i <= n; i++ {
		b1, b2 = b1+b2, b1
	}
	return b1
}

/*
func climbStairs(n int) int {
	ways, b1, b2 := 1, 1, 1
	for i := 2; i <= n; i++ {
		ways, b1, b2 = b1+b2, b1+b2, b1
	}
	return ways
}
*/

// 标准动态规划解法
func ClimbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

package problems

/* greedy O(n)
贪心算法，子问题最优解递推到最终问题最优解。
只要后一天的价格高于当天，那当天就买入，第二天卖出。
遍历数组求每两天的最大利润，最后将利润相加
*/
func MaxProfit(prices []int) int {
	profits := 0
	for i := range prices {
		if i == 0 {
			continue
		}
		if prices[i] > prices[i-1] {
			profits += prices[i] - prices[i-1]
		}
	}
	return profits
}

/*
贪心算法每次只考虑当前子问题的最优解，
而动态规划会记录每个分支的状态，
本例贪心算法没有问题，当需要手续费或者由交易限制时，
就需要考虑动态规划
*/

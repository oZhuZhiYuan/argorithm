package other

/*
动态规划实例问题
*/

/*
chess := [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, -1, 0, 0, 0, -1, 0},
	{0, 0, 0, 0, -1, 0, 0, 0},
	{-1, 0, -1, 0, 0, -1, 0, 0},
	{0, 0, -1, 0, 0, 0, 0, 0},
	{0, 0, 0, -1, -1, 0, -1, 0},
	{0, -1, 0, 0, 0, -1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}
*/

func NumPath(chess [][]int) int {
	hi := len(chess) - 1
	wi := len(chess[0]) - 1
	// 初始化
	for i := 0; i < hi; i++ {
		chess[i][wi] = 1
	}
	for j := 0; j < wi; j++ {
		chess[hi][j] = 1
	}
	// 状态方程 chess[i][j] = chess[i+1][j] + chess[i][j+1]
	// 逐行遍历
	for i := hi - 1; i >= 0; i-- {
		for j := wi - 1; j >= 0; j-- {
			if chess[i][j] == -1 {
				chess[i][j] = 0
				continue
			}
			chess[i][j] = chess[i+1][j] + chess[i][j+1]
		}
	}

	return chess[0][0]
}

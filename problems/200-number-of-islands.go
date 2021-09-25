package problems

/*
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

*/

/*
解法一 遍历+dfs染色
如果发现陆地，将周围陆地都炸平，计数器加1
注意: 切片是引用类型
*/
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	num := 0
	// 遍历
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				// 如果发现陆地，将周围陆地都炸平，计数器加1
				floodFillDfs(i, j, grid)
				num++
			}
		}
	}
	return num
}

func floodFillDfs(x, y int, grid [][]byte) {
	// isValid,注意先后顺序
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	// 炸平
	(grid)[x][y] = '0'
	// 递归遍历上下左右
	floodFillDfs(x-1, y, grid)
	floodFillDfs(x+1, y, grid)
	floodFillDfs(x, y-1, grid)
	floodFillDfs(x, y+1, grid)
}

/*
解法二 并查集

*/

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
解法：如果发现陆地，将周围陆地都炸平，计数器加1
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
解法：
1.创建disjoint sets, 初始状态下每个陆地的root都是自己
2.遍历二维数组
3.如果发现陆地，就遍历上下左右，发现陆地便执行union
4.最后返回disjoint set的数量
*/

type UnionFind struct {
	grid   [][]byte
	parent []int
	rank   []int
	count  int
}

func ufConstructor(grid [][]byte) UnionFind {
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				count++
			}
		}
	}

	return UnionFind{grid: grid, parent: parent, rank: rank, count: count}
}

func (uf *UnionFind) find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFind) union(x, y int) {
	rootx := uf.find(x)
	rooty := uf.find(y)
	if rootx != rooty {
		if uf.rank[rootx] > uf.rank[rooty] {
			uf.parent[rooty] = rootx
		} else if uf.rank[rootx] < uf.rank[rooty] {
			uf.parent[rootx] = rooty
		} else {
			uf.parent[rooty] = rootx
			uf.rank[rootx]++
		}
		// 每次合并，减去一个集合计数
		uf.count--
	}
}

func NumIslands1(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// 初始化并查集数据结构
	uf := ufConstructor(grid)
	m, n := len(grid), len(grid[0])
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 遍历
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			for _, dr := range directions {
				x, y := i+dr[0], j+dr[1]
				if x >= 0 && x < m && y >= 0 && y < m && uf.grid[x][y] == '1' {
					uf.union(i*n+j, x*n+y)
				}
			}

		}
	}
	return uf.count
}

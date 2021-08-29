package problems

/*
DFS & 位运算
注意位运算的优先级
*/
func TotalNQueens(n int) int {
	var count int
	dfs(&count, n, 0, 0, 0, 0)
	return count
}

func dfs(count *int, n, row, cols, pie, na int) {
	// recursion terminator
	if row >= n {
		*count++
		return
	}

	/*
		prune:
		calculate avilable postion(bit) in this level and set to "1", irrelevent bits set to "0",
		&(1<<n-1) is the trick ignore the irrelevent bits.
		eg. 0000 0000 0000 0000 0000 0000 0000 1110 (n=4)
	*/
	bitsForQ := ^(cols | pie | na) & (1<<n - 1)

	for bitsForQ > 0 {
		// get the lowest "1", which stands for currnet avilable postion, eg. 1110 -> 0010 (n=4)
		p := bitsForQ & -bitsForQ
		dfs(count, n, row+1, p|cols, (p|pie)<<1, (p|na)>>1)
		// set the lowest "1" to "0", eg. 1110 -> 1100 (n=4)
		bitsForQ &= (bitsForQ - 1)
	}
}

/*
eg.  n=4  the second round recursion: (first one bitsForQ is 1111)

level0 ..Q.  n = 4, row = 0,                                     bitsForQ = 1110, p = 0010
level1 Q...  n = 4, row = 1, cols = 0010, pie = 0100, na = 0001, bitsForQ = 1000, p = 1000
level2 ...Q  n = 4, row = 2, cols = 1010, pie = 1000, na = 0100, bitsForQ = 0001, p = 0001
level3 .Q..  n = 4, row = 3, cols = 1011, pie = 0010, na = 0010, bitsForQ = 0100, p = 0100
level4       n = 4, row = 4, ... *count++

*/

package problems

/*
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

广度优先搜索 BFS

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	var queue []*TreeNode
	queue = append(queue, root)

	// 一次遍历一层
	for len(queue) > 0 {
		len := len(queue)
		var tmp []int
		for _, node := range queue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, tmp)
		queue = queue[len:]
	}
	return ret
}

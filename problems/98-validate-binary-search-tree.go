package problems

// https://leetcode-cn.com/problems/validate-binary-search-tree/

/*
In-order traversal the bst, every pre node's val should less than current node's val.
*/
func recur(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}
	if !recur(root.Left, pre) {
		return false
	}
	if *pre >= root.Val {
		return false
	}
	*pre = root.Val
	return recur(root.Right, pre)
}

func IsValidBST(root *TreeNode) bool {
	// init pre as mininum of type int
	pre := ^int(^uint(0) >> 1)
	return recur(root, &pre)
}

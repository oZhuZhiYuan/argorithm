package problems

/**
二叉树dfs + 剪枝

题干：
https://leetcode-cn.com/problems/generate-parentheses/

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

输入：n = 1
输出：["()"]


题解：
            root
		    /  \
		   (    )
		  / \   / \
         (   ) (   )

括号只有左右小括号两种，可以等价于深度2n跳的二叉树DFS,每次到底便是一个结果，
判断每个结果是不是有效的括号集合，时间复杂度O(2^2n)。

需要通过剪枝来加快：
1、右括号开头剪掉;
2、左括号和右括号的配额是n;
3、每次递归时，右括号的数量应该比左括号少。

*/

func GenerateParenthesis(n int) []string {
	var pas []string
	var ret []byte // 使用[]byte 比string 更节省空间
	genpas(&pas, ret, 0, 0, n)
	return pas
}

func genpas(pas *[]string, ret []byte, left, right, n int) {
	if left == n && right == n {
		*pas = append(*pas, string(ret))
		return
	}
	if left < n {
		genpas(pas, append(ret, '('), left+1, right, n)
	}
	if right < left && right < n {
		genpas(pas, append(ret, ')'), left, right+1, n)
	}
}

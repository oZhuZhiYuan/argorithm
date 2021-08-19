package testcase

import (
	pro "argo/problems"
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	var root *pro.TreeNode
	if pro.IsValidBST(root) {
		fmt.Println("true")
	}
}

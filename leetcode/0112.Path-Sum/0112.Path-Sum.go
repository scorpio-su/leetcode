package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return hasPathSum(root.Right, sum-root.Val) || hasPathSum(root.Left, sum-root.Val)
}

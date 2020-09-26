package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

// BFS ways
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return issometree(root.Left, root.Right)
}

func issometree(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}
	if l.Val != r.Val {
		return false
	}
	return issometree(l.Left, r.Right) && issometree(l.Right, r.Left)
}

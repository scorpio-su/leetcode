package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := minDepth(root.Left), minDepth(root.Right)
	if l == 0 || r == 0 {
		return l + r + 1
	}
	return 1 + min(l, r)
}

// use own funtion do min
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

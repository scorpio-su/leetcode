package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth1(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth1(root.Left) + 1
	}
	return min(minDepth1(root.Left), minDepth1(root.Right)) + 1
}

// use own funtion do min
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

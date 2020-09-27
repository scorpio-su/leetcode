package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		l, r := maxDepth(root.Left), maxDepth(root.Right)
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
}

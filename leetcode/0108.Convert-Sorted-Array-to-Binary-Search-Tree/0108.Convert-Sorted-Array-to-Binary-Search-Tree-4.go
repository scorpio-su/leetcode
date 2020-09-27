package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

// do cut the half and use helper
func sortedArrayToBST4(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	mid := l / 2
	return &TreeNode{Val: nums[mid], Left: sortedArrayToBST4(nums[:mid]), Right: sortedArrayToBST4(nums[mid+1:])}
}

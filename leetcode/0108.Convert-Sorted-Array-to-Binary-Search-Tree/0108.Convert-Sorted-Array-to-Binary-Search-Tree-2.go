package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

// do cut the half
func sortedArrayToBST2(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[l/2]
	root.Left = sortedArrayToBST2(nums[:l/2])
	root.Right = sortedArrayToBST2(nums[l/2+1:])
	return root
}

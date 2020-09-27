package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

// do cut the half and use helper
func sortedArrayToBST3(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = helper(nums, l, mid-1) //mid have had
	root.Right = helper(nums, mid+1, r)
	return root
}

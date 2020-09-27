package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

// do cut the half
func sortedArrayToBST1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST1(nums[:len(nums)/2]),
		Right: sortedArrayToBST1(nums[len(nums)/2+1:]),
	}
}

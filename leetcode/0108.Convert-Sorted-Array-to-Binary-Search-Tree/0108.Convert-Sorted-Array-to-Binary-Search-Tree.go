package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

// do cut the half
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	n /= 2
	return &TreeNode{
		Val:   nums[n],
		Left:  sortedArrayToBST(nums[:n]),
		Right: sortedArrayToBST(nums[n+1:]),
	}

}

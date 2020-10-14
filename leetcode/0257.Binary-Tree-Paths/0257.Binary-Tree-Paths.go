package leetcode

import (
	"strconv"

	"github.com/scorpio-su/leetcode/structures"
)

type TreeNode = structures.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	dfs(root, &ans, "")
	return ans
}

func dfs(root *TreeNode, s *[]string, s1 string) {
	if root == nil {
		return
	}
	s1 += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*s = append(*s, s1)
	}
	s1 += "->"
	// left
	dfs(root.Left, s, s1)
	// right
	dfs(root.Right, s, s1)
}

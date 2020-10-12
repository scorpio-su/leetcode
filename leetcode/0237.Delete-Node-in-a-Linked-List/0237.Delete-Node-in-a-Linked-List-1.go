package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type ListNode = structures.ListNode

func deleteNode1(node *ListNode) {
	*node = *node.Next
}

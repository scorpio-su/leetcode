package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type ListNode = structures.ListNode

func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}

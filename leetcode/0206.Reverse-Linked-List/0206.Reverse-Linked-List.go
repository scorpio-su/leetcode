package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var ans *ListNode
	for head != nil {
		next := head.Next
		head.Next = ans
		ans = head
		head = next
	}
	return ans
}

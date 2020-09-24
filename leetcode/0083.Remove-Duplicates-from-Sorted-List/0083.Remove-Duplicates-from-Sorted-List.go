package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

// use own package
type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.val == cur.Next.val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

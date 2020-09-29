package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type ListNode = structures.ListNode

func hasCycle(head *ListNode) bool {
	for fast, slow := head, head; fast != nil && fast.Next != nil; {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{Val: l1.Val + l2.Val}
	durr := ans
	l1 = l1.Next
	l2 = l2.Next
	next := new(ListNode)
	for l1 != nil || l2 != nil || durr.Val > 0 {
		if durr.Val > 9 { //deal with Val>=10
			next.Val++
			durr.Val -= 10
		}

		if next.Val == 0 && (l1 == nil || l2 == nil) {
			if l1 != nil {
				durr.Next = l1
			} else {
				durr.Next = l2
			}
			break
		}

		if l1 != nil {
			next.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			next.Val += l2.Val
			l2 = l2.Next
		}

		if next.Val > 0 || l1 != nil || l2 != nil {
			durr.Next = next
			durr = next
			next = new(ListNode)
		}
	}
	return ans
}

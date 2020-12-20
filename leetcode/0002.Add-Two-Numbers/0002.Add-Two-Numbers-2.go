package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	durr := ans
	crr := 0
	for {
		if l1 == nil && l2 == nil && crr == 0 {
			break
		}
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + crr
		durr.Next = &ListNode{Val: sum % 10}
		crr = sum / 10
		durr = durr.Next
	}
	return ans.Next
}

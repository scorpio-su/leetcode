package leetcode

import (
	"github.com/scorpio-su/leetcode/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := ListNode{0, nil}
	durr := &ans
	crr := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}

		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}

		sum := x + y + crr
		crr = 0
		if sum >= 10 {
			crr = sum / 10
			sum %= 10
		}

		durr.Next = &ListNode{sum, nil}
		durr = durr.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	if crr != 0 {
		durr.Next = &ListNode{crr, nil}
	}
	return ans.Next
}

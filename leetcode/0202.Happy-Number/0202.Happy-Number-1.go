package leetcode

// two and one go

func isHappy(n int) bool {
	slow, fast := gethappy(n), gethappy(gethappy(n))
	for {
		if fast == 1 {
			return true
		} else if slow == fast {
			return false
		} else {
			slow = gethappy(slow)
			fast = gethappy(gethappy(fast))
		}
	}
}

func gethappy(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

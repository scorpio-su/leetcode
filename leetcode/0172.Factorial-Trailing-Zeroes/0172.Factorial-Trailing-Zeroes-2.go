package leetcode

func trailingZeroes2(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes2(n/5)
}

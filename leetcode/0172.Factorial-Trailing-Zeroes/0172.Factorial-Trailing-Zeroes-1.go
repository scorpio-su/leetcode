package leetcode

func trailingZeroes1(n int) int {
	ans := 0
	for n/5 > 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}

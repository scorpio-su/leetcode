package leetcode

func isPowerOfTwo1(n int) bool {
	return n&(n-1) == 0 && n > 0
}

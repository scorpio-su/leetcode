package leetcode

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n%2 == 1 {
		return false
	}
	return isPowerOfTwo(n / 2)
}

package leetcode

func isPowerOfTwo2(n int) bool {
	if n == 0 {
		return false
	}
	i := 1
	for i < n {
		n <<= 1
	}
	return n == 1
}

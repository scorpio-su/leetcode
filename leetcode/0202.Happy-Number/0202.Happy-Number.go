package leetcode

func isHappy(n int) bool {
	m := make(map[int]bool)
	for !m[n] && n != 1 {
		m[n] = true
		n = gethappy(n)
	}
	return n == 1
}

func gethappy(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

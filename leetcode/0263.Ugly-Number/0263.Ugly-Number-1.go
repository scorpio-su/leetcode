package leetcode

func isUgly1(num int) bool {
	if num == 1 {
		return true
	}
	if num <= 0 {
		return false
	}
	for _, i := range [3]int{2, 3, 5} {
		for num%i == 0 {
			num /= i
		}
	}
	return num == 1
}

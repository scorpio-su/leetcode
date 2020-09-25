package leetcode

func plusOne(digits []int) []int {
	extra := 1
	for i := len(digits) - 1; i >= 0; i-- {
		extra += digits[i]
		digits[i] = extra % 10
		extra /= 10
	}
	if extra != 0 {
		digits = append([]int{extra}, digits...)
	}
	return digits
}

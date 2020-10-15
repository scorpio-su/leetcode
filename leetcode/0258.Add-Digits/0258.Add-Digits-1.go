package leetcode

func addDigits1(num int) int {
	if num < 10 {
		return num
	}
	mod := num % 9
	if mod == 0 {
		return 9
	}
	return mod
}

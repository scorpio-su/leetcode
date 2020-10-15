package leetcode

func addDigits(num int) int {
	var digit int
	for num > 9 {
		digit = 0
		for num != 0 {
			digit += num % 10
			num /= 10
		}
		num = digit
	}
	return num
}

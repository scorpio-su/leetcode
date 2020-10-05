package leetcode

func titleToNumber(s string) int {
	num, ans := 0, 0
	for _, c := range s {
		num = int(c - 'A' + 1)
		ans = ans*26 + num
	}
	return ans
}

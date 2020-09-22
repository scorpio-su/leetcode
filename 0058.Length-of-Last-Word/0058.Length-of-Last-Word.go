package leetcode

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && ans == 0 {
			continue
		}
		if s[i] == ' ' && ans != 0 {
			break
		}
		ans++
	}
	return ans
}

package leetcode

func lengthOfLastWord1(s string) int {
	ans, i := 0, len(s)-1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		ans++
	}
	return ans
}

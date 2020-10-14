package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := make(map[rune]int)
	for _, r := range s {
		n[r]++
	}
	for _, r := range t {
		n[r]--
	}
	for k := range n {
		if n[k] != 0 {
			return false
		}
	}
	return true
}

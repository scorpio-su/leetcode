package leetcode

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := [26]int{}
	for _, r := range s {
		n[r-'a']++
	}
	for _, r := range t {
		n[r-'a']--
	}
	for i := 0; i < 26; i++ {
		if n[i] != 0 {
			return false
		}
	}
	return true
}

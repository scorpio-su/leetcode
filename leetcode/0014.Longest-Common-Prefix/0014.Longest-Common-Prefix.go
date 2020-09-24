package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for _, s := range strs {
		//測試兩者的string是否一樣
		for !strings.HasPrefix(s, pre) {
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}

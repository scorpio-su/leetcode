package leetcode

func isValid(s string) bool {
	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	} else if len(s)%2 == 1 {
		return false
	}
	data := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) == 0 || c != data[stack[len(stack)-1]] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

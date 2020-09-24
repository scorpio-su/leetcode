package leetcode

func addBinary(a string, b string) string {
	i, j, l := len(a)-1, len(b)-1, []byte{}
	var carry byte
	for i >= 0 || j >= 0 || carry == 1 {
		if i >= 0 {
			carry += a[i] - '0'
			i--
		}
		if j >= 0 {
			carry += b[j] - '0'
			j--
		}
		l = append(l, carry%2+'0')
		carry /= 2
	}
	return string(reverse(l))
}

func reverse(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}

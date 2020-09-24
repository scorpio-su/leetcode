package leetcode

func mySqrt1(x int) int {
	if x == 1 {
		return 1
	}
	for i, j := 0, x; ; {
		m := i + (j-i)/2
		if m*m > x {
			j = m + 1
		} else if (m+1)*(m+1) > x {
			return m
		} else {
			i = m - 1
		}
	}
}

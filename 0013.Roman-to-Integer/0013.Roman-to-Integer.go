package leetcode

var dir = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	length := len(s)
	total, num, last := 0, 0, 0
	for i := 0; i < length; i++ {
		num = dir[s[length-i-1:length-i]]
		if last > num {
			total -= num
		} else {
			total += num
		}
		last = num
	}
	return total
}

package leetcode

func hammingWeight1(num uint32) int {
	var cout uint32
	for num > 0 {
		cout += num & 1
		num >>= 1
	}
	return int(cout)
}

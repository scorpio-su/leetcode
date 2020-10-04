package leetcode

func majorityElement(nums []int) int {
	nu := len(nums)
	n := make(map[int]int, nu)
	nu /= 2
	for _, q := range nums {
		n[q]++
		if n[q] > nu {
			return q
		}
	}
	return -1
}

package leetcode

func containsDuplicate(nums []int) bool {
	n := make(map[int]int, len(nums))
	for _, op := range nums {
		if n[op] != 0 {
			return true
		}
		n[op]++
	}
	return false
}

package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	n := make(map[int]int, len(nums))
	for i, v := range nums {
		if q, ok := n[v]; ok {
			if i-q <= k {
				return true
			}
		}
		n[v] = i
	}
	return false
}

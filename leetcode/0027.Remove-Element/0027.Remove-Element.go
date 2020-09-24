package leetcode

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
	}

	return l
}

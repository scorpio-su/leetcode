package leetcode

func rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	reverse(nums, 0, length-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

func reverse(nums []int, st, en int) {
	for st < en {
		nums[en], nums[st] = nums[st], nums[en]
		en--
		st++
	}
}

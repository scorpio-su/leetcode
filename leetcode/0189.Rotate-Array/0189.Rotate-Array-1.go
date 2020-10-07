package leetcode

func rotate1(nums []int, k int) {
	length := len(nums)
	k %= length
	if len(nums) == 0 || k <= 0 {
		return
	}
	tem := make([]int, length)
	copy(tem, nums[length-k:])
	copy(tem[k:], nums)
	copy(nums, tem)
}

package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max1 := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max1 {
			max1 = nums[i]
		}
	}
	return max1
}

package leetcode

func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans, tem := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if tem >= 0 {
			tem += nums[i]
		} else {
			tem = nums[i]
		}
		ans = max(ans, tem)
	}
	return ans
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

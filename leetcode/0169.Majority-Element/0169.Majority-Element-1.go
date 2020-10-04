package leetcode

import "sort"

func majorityElement1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// use sort funtion to do
	sort.Ints(nums)
	return nums[len(nums)/2]
}

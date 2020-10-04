package leetcode

func majorityElement2(nums []int) int {
	ans, s := 0, 0
	for _, num := range nums {
		if ans == 0 {
			s = num
		}
		if num == s {
			ans++
		} else {
			ans--
		}
	}
	return s
}

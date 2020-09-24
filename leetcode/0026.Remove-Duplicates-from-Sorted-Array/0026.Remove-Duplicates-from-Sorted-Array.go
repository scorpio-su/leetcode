package leetcode

func removeDuplicates(num []int) int {
	if len(num) == 0 {
		return 0
	}
	rat := 0
	for i := 0; i < len(num); i++ {
		if num[i] != num[rat] {
			rat++
			num[rat] = num[i]
		}
	}
	return rat + 1
}

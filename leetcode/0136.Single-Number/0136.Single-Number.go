package leetcode

func singleNumber(nums []int) int {
	x := 0
	for _, num := range nums {
		x ^= num
	}
	return x
}

// use xor
/*
0 0 | 0
0 1 | 1
1 0 | 1
1 1| 0
*/

package leetcode

func getRow1(rowIndex int) []int {
	result := make([]int, 0, rowIndex+1)
	for ; rowIndex >= 0; rowIndex-- {
		result = append(result, 1)
		for j := len(result) - 2; j > 0; j-- {
			result[j] += result[j-1]
		}
	}
	return result
}

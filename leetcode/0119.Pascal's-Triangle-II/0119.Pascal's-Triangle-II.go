package leetcode

func getRow(rowIndex int) []int {
	result := make([][]int, rowIndex+2)
	for i := 0; i < rowIndex+1; i++ {
		a := make([]int, i+1)
		a[0], a[i] = 1, 1
		for j := 1; j < i; j++ {
			a[j] = result[i-1][j] + result[i-1][j-1]
		}
		result[i] = a
	}
	return result[rowIndex]
}

package leetcode

func generate1(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		a := make([]int, i+1)
		a[0], a[i] = 1, 1
		for j := 1; j < i; j++ {
			a[j] = result[i-1][j] + result[i-1][j-1]
		}
		result[i] = a
	}
	return result
}

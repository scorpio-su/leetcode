package leetcode

func generate(numRows int) [][]int {
	var result [][]int
	for i := 1; i <= numRows; i++ {
		a := make([]int, i)
		a[0], a[len(a)-1] = 1, 1
		for j := 1; j < len(a)-1; j++ {
			a[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, a)
	}
	return result
}

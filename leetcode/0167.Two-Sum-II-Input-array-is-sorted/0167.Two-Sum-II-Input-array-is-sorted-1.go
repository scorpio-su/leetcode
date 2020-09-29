package leetcode

func twoSum1(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	sums := numbers[i] + numbers[j]
	for sums != target {
		if sums > target {
			j--
		} else if sums < target {
			i++
		}
		sums = numbers[i] + numbers[j]
	}
	return []int{i + 1, j + 1}
}

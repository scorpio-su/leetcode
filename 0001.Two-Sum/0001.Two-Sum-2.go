package leetcode

func twoSum2(num []int, target int) []int {
	m := make(map[int]int, len(num))
	for i, n := range num {
		if c, ok := m[n]; ok {
			return []int{c, i}
		}
		m[target-n] = i
	}
	return nil
}

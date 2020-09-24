package leetcode

func twoSum(num []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(num); i++ {
		anthor := target - num[i]
		if _, ok := m[anthor]; ok {
			return []int{m[anthor], i}
		}
		m[num[i]] = i
	}
	return nil
}

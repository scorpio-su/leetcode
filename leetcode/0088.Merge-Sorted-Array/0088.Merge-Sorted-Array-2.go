package leetcode

func merge2(nums1 []int, m int, nums2 []int, n int) {
	// use var k beck
	i, j, k := m, n, m+n

	for j > 0 && k > 0 {
		if i > 0 && nums1[i-1] > nums2[j-1] {
			nums1[k-1] = nums1[i-1]
			i--
		} else {
			nums1[k-1] = nums2[j-1]
			j--
		}
		k--
	}
}

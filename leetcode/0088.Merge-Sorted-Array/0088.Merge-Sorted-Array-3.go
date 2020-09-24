package leetcode

func merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	// submit is unnecessary
	//the code, test don't need
	// for i >= 0 {
	//     nums1[k] = nums1[i]
	//     i--
	//     k--
	// }
}

package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < len(nums1) && n > j {
		if nums1[i] > nums2[j] || i >= m {
			nums1 = append(nums1[:i+1], nums1[i:len(nums1)-1]...)
			nums1[i] = nums2[j]
			j++
			m++
		}
		i++
	}
}

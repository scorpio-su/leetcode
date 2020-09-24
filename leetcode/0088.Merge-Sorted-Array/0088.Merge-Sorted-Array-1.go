package leetcode

func merge1(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		//sort from behind
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

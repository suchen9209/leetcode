package main

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m+n && j < n {
		if nums2[j] < nums1[i] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++
		} else {
			i++
		}
	}
	for j < n {
		nums1[m+j] = nums2[j]
		j++
	}
}

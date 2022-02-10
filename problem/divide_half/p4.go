package main

import "fmt"

func main() {
	//nums1 := []int{1,3,4,5,6,8,10,11,15,17}
	//nums2 := []int{2,2,3,5,6,7,8,9,20,45,50,55,60}
	nums1 := []int{1}
	nums2 := []int{2, 3, 4, 5, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func getk(nums1 []int, nums2 []int, k int) int {
	if k == 1 {
		if len(nums1) == 0 {
			return nums2[0]
		} else if len(nums2) == 0 {
			return nums1[0]
		}
		if nums1[0] > nums2[0] {
			return nums2[0]
		} else {
			return nums1[0]
		}
	}

	half := k/2 - 1
	if len(nums2) > half && (len(nums1) <= half || nums1[half] > nums2[half]) {
		return getk(nums1, nums2[k/2:], k-k/2)
	} else {
		return getk(nums1[k/2:], nums2, k-k/2)
	}

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//总和为奇数，直接返回第(len(nums1) + len(nums2))/ 2大的
	//总和为偶数,返回第(len(nums1) + len(nums2))/ 2 和 (len(nums1) + len(nums2))/ 2 - 1大的数值的中间值
	midPos := (len(nums1)+len(nums2))/2 + 1
	if (len(nums1)+len(nums2))%2 == 0 {
		return (float64(getk(nums1, nums2, midPos)) + float64(getk(nums1, nums2, midPos-1))) / 2
	} else {
		return float64(getk(nums1, nums2, midPos))
	}

}

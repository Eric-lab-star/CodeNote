package main

import "fmt"

func main() {
	fmt.Println("Leetcode")
	nums1 := []int{0}
	m := 0
	n := 1
	nums2 := []int{1}
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; i < m+n; i, j = i+1, j+1 {
		nums1[i] = nums2[j]
	}
	fmt.Println(nums1)
	for i := m; i < len(nums1); i++ {
		if i != 0 && nums1[i] < nums1[i-1] {
			for j := i; j != 0 && nums1[j] < nums1[j-1]; j-- {
				nums1[j], nums1[j-1] = nums1[j-1], nums1[j]
			}
		}

	}
}

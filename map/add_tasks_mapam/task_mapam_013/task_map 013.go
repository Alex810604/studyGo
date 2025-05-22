// 13. Проверить, является ли массив подмножеством другого
// Даны два массива. Проверьте, все ли элементы первого массива присутствуют во втором.
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{2, 4, 6, 21, 8, 9, 10, 12, 14, 16, 18, 19, 20}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(subsetArray(nums1, nums2))
	fmt.Println(subsetArray1(nums1, nums2))
}
func subsetArray(nums1, nums2 []int) bool {
	m2 := make(map[int]struct{})
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}
	fmt.Println(m2)
	for _, v := range nums1 {
		if _, ok := m2[v]; !ok {
			return false
		}
	}
	return true
}

// или булевый map
func subsetArray1(nums1, nums2 []int) bool {
	m2 := make(map[int]bool)
	for _, v := range nums2 {
		m2[v] = true
	}
	fmt.Println(m2)
	for _, v := range nums1 {
		if !m2[v] {
			return false
		}
	}
	return true // true	false
}

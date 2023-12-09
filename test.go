package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
loop1:
	for {
		fmt.Println("loop1")
	Loop2:
		for k, v := range nums {
			if v == val {
				nums = append(nums[:k], nums[k+1:]...)
				break Loop2
			}
			if (k + 1) == len(nums) {
				break loop1
			}
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func merge(nums1 []int, n int, nums2 []int, m int) []int {
	result := make([]int, 0, n+m)
	for _, v := range nums1 {
		if v != 0 {
			result = append(result, v)
		}
	}
	for _, v := range nums2 {
		if v != 0 {
			result = append(result, v)
		}
	}
	nums1 = result
	sort.Ints(nums1)
	fmt.Println(nums1)
	return nums1
}

func main() {
	num := []int{3, 2, 2, 3}
	removeElement(num, 3)
	sad := []int{}
	fmt.Println(len(sad))
	for k := range sad {
		fmt.Println(len(sad), k)
	}
}

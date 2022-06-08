package leetcode

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	mp := make(map[int]bool)
	for _, n := range nums {
		if _, ok := mp[n]; !ok {
			mp[n] = true
		} else {
			return true
		}
	}
	return false
}

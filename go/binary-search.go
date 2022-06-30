package leetcode

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	for left <= right {
		m := left + (right-left)/2

		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return -1
}

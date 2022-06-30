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

// Recursive
func searchRecursive(nums []int, target int) int {
	return binarySearchRecursive(nums, 0, len(nums)-1, target)
}

func binarySearchRecursive(nums []int, left int, right int, target int) int {
	if left <= right {
		m := left + (right-left)/2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			return binarySearchRecursive(nums, left, m-1, target)
		}
		return binarySearchRecursive(nums, m+1, right, target)
	}
	return -1
}

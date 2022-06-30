package leetcode

// Iterative
func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l int, r int, target int) int {
	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

// Recursive
func searchInsertRecursive(nums []int, target int) int {
	return binarySearchRecursive(nums, 0, len(nums)-1, target)
}

func binarySearchRecursive(nums []int, l int, r int, target int) int {
	if l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			return binarySearchRecursive(nums, l, m-1, target)
		}
		return binarySearchRecursive(nums, m+1, r, target)
	}
	return l
}

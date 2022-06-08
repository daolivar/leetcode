package leetcode

// O(n) time
// O(n) space
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, num := range nums {
		if _, ok := mp[target-num]; ok {
			return []int{mp[target-num], i}
		}
		mp[num] = i
	}
	return []int{}
}

// Brute Force
// func twoSum(nums []int, target int) []int {
// 	res := make([]int, 2)
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums) && i != j; j++ {
// 			if nums[i]+nums[j] == target {
// 				res[0] = i
// 				res[1] = j
// 				return res
// 			}
// 		}
// 	}
// 	return res
// }

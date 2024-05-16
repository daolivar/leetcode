/**
 * Solution class for LeetCode problem 1. Two Sum
 */
public class Solution {

    /**
     * Finds two numbers in the given array that add up to the target sum.
     * 
     * Uses a HashMap to store the differences between the target and the current element.
     * 
     * @param nums an array of integers
     * @param target the target sum
     * @return an array containing the indices of the two numbers whose sum equals the target, or null if no such pair exists
     */
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> diffs = new HashMap<>(); // Map to store differences between target and current element
        for (int idx = 0; idx < nums.length; idx++) {
            int num = nums[idx]; // Current number
            int comp = target - num; // Complement of the current number
            if (diffs.containsKey(comp)) { // If complement exists in the map
                return new int[]{diffs.get(comp), idx}; // Return indices of the two numbers
            }
            diffs.put(num, idx); // Store current number and its index in the map
        }
        return null; // No solution found
    }
}

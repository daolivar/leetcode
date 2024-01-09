/**
 * Solution to find two numbers in an array that add up to a specific target.
 */
class Solution {
    /**
     * Finds two numbers in the given array that add up to the target.
     * 
     * Time Complexity: O(n), where 'n' is the length of the input array nums. In
     * the worst case, the algorithm iterates through the entire array once.
     * 
     * Space Complexity: O(n), where 'n' is the length of the input array nums. In
     * the worst case, the algorithm may store all elements of the array in the map
     * along with their indices.
     * 
     * @param nums   The array of integers.
     * @param target The target sum.
     * @return An array containing the indices of the two numbers that add up to the
     *         target. If no such pair is found, an empty array is returned.
     */
    public int[] twoSum(int[] nums, int target) {
        // Map to store the difference of each element and its index
        Map<Integer, Integer> map = new HashMap<>();

        // Iterate through the array
        for (int i = 0; i < nums.length; i++) {
            int difference = target - nums[i];

            // Check if the difference is present in the map
            if (map.containsKey(difference)) {
                // Return the indices of the two numbers
                return new int[] { map.get(difference), i };
            }

            // If difference not found, store the current element and its index in the map
            map.put(nums[i], i);
        }

        // If no such pair is found, return an empty array
        return new int[] {};
    }
}

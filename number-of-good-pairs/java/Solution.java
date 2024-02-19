/**
 * Solution class provides a method to count the number of good pairs in an array.
 */
class Solution {
    /**
     * Counts the number of good pairs in the given array.
     * 
     * A good pair (i, j) is a pair where i < j and nums[i] == nums[j].
     * 
     * Time Complexity: O(n), where n is the length of the nums array. We iterate
     * through the nums array once.
     * 
     * Space Complexity: O(n), where n is the length of the nums array. We use a
     * HashMap to store the frequency of each number in the array.
     * 
     * @param nums An array of integers.
     * @return The number of good pairs in the array.
     */
    public int numIdenticalPairs(int[] nums) {
        int goodPairs = 0;
        HashMap<Integer, Integer> frequencyMap = new HashMap<>();

        for (int num : nums) {
            // Increment the count of good pairs by the frequency of the current number
            goodPairs += frequencyMap.getOrDefault(num, 0);
            // Update the frequency of the current number in the map
            frequencyMap.put(num, frequencyMap.getOrDefault(num, 0) + 1);
        }

        return goodPairs;
    }
}

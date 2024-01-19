/**
 * Solution class provides methods to find a non-minimum or non-maximum element
 * in an array.
 */
class Solution {
    /**
     * Finds a non-minimum or non-maximum element in the given array.
     * 
     * Time Complexity: O(n), where 'n' is the length of the input array nums. The
     * algorithm iterates through the entire array once.
     * 
     * Space Complexity: O(1), as the space used is constant, independent of the
     * input size.
     * 
     * @param nums The array of integers.
     * @return A non-minimum or non-maximum element in the array. If no such element
     *         is found, returns -1.
     */
    public int findNonMinOrMax(int[] nums) {
        // Find the minimum and maximum values in the array
        int[] minAndMax = findMinAndMax(nums);
        int min = minAndMax[0], max = minAndMax[1];

        // Iterate through the array to find a non-minimum and non-maximum element
        for (int i : nums) {
            if (i != min && i != max) {
                return i;
            }
        }

        // If no such element is found, return -1
        return -1;
    }

    /**
     * Finds the minimum and maximum values in the given array.
     * 
     * Time Complexity: O(n), where 'n' is the length of the input array nums. The
     * algorithm iterates through the entire array once.
     * 
     * Space Complexity: O(1), as the space used is constant, independent of the
     * input size.
     * 
     * @param nums The array of integers.
     * @return An array containing the minimum and maximum values in the input
     *         array.
     */
    public static int[] findMinAndMax(int[] nums) {
        int max = Integer.MIN_VALUE;
        int min = Integer.MAX_VALUE;

        // Iterate through the array to find the minimum and maximum values
        for (int i : nums) {
            max = Math.max(max, i);
            min = Math.min(min, i);
        }

        // Return an array containing the minimum and maximum values
        return new int[] { min, max };
    }
}

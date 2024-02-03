/**
 * Solution that determines whether an array of integers contains duplicate
 * values. Uses a HashSet to efficiently track previously encountered values.
 */
class Solution {
    /**
     * Checks for duplicate values in the given array.
     * 
     * Time Complexity: O(n) where n is the length of the input array nums. This
     * is because, in the worst case, the function iterates through the entire array
     * once, and each set operation (contains and add) takes constant time on
     * average.
     * 
     * Space Complexity: O(n) in the worst case, where all elements in the array
     * are unique, the set set will contain all n elements. Therefore, the space
     * required is linearly proportional to the size of the input array.
     * 
     * @param nums The array of integer values to be examined.
     * @return {@code true} if duplicates are found, {@code false} otherwise.
     */
    public boolean containsDuplicate(int[] nums) {
        // Hashset to store unique values encountered so far
        Set<Integer> set = new HashSet<>();

        // Iterate through the array
        for (int n : nums) {
            // If the value is already in the set, a duplicate is found
            if (set.contains(n)) {
                return true;
            }
            // Otherwise, add the value to the set
            set.add(n);
        }

        // No duplicate found in the array
        return false;
    }
}

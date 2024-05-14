/**
 * Solution classes for Leetcode problem 217. Contains Duplicate
 */
public class Solution {

    /**
     * Checks whether the given array contains any duplicate elements.
     * 
     * This method utilizes a HashSet data structure to efficiently identify duplicate elements.
     * 
     * @param nums an array of integers to check for duplicates
     * @return true if the array contains any duplicate elements, otherwise false
     */
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> set = new HashSet<>(); // Create a HashSet to store unique elements
        for (int i : nums) { // Iterate through each element in the input array
            if (set.contains(i)) { // Check if the HashSet already contains the current element
                return true; // If a duplicate is found, return true
            }
            set.add(i); // Add the element to the HashSet to keep track of seen elements
        }
        return false; // If no duplicates are found, return false
    }
}


public class SolutionSort {

    /**
     * Checks whether the given array contains any duplicate elements.
     * 
     * This method sorts the array and then iterates through it, comparing each element with the next one.
     * If any adjacent elements are found to be equal, it indicates the presence of duplicates.
     * 
     * @param nums an array of integers to check for duplicates
     * @return true if the array contains any duplicate elements, otherwise false
     */
    public boolean containsDuplicate(int[] nums) {
        Arrays.sort(nums); // Sort the array in ascending order
        for (int i = 0; i < nums.length - 1; i++) { // Iterate through the array, comparing each element with the next one
            if (nums[i] == nums[i + 1]) { // If any adjacent elements are found to be equal
                return true; // Return true indicating the presence of duplicates
            }
        }
        return false; // If no duplicates are found, return false
    }
}

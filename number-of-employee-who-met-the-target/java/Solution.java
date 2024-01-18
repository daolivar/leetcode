/**
 * Solution class provides a method to count the number of employees who met or
 * exceeded a target number of hours.
 */
class Solution {
    /**
     * Counts the number of employees who met or exceeded the target number of
     * hours.
     * 
     * Time Complexity: O(n), where 'n' is the length of the input array hours. The
     * algorithm iterates through the entire array once.
     * 
     * Space Complexity: O(1), as the space used is constant, independent of the
     * input size.
     * 
     * @param hours  An array representing the hours worked by each employee.
     * @param target The target number of hours.
     * @return The count of employees who met or exceeded the target number of
     *         hours.
     */
    public int numberOfEmployeesWhoMetTarget(int[] hours, int target) {
        // Count tracks the number of employees meeting the target.
        int count = 0;

        // Iterate through the array of hours.
        for (int i : hours) {
            // Check if the hours worked by the employee is greater than or equal to the
            // target.
            if (i >= target) {
                // Increment the count.
                count++;
            }
        }

        // Return the final count.
        return count;
    }
}

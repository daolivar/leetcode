/**
 * Solution class provides a method to find the smallest even multiple of a given integer.
 */
class Solution {
    /**
     * Finds the smallest even multiple of the given integer.
     * 
     * Time Complexity: O(1), as the operation involves a simple conditional check and multiplication, both constant time.
     * 
     * Space Complexity: O(1), as there is no additional space used that grows with the input.
     * 
     * @param n The input integer.
     * @return The smallest even multiple of the input integer. If the input is already even, returns the input itself.
     */
    public int smallestEvenMultiple(int n) {
        return n % 2 == 0 ? n : 2 * n;
    }
}

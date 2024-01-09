/**
 * Solution to calculate the number of distinct ways to climb stairs using
 * dynamic programming.
 */
public class Solution {
    /**
     * Calculates the number of distinct ways to climb stairs.
     * 
     * Time Complexity: O(n) This is because the function uses a bottom-up dynamic
     * programming approach (tabulation) to fill the array dp from the base cases (0
     * and 1) up to n. The loop iterates n times, and at each iteration, a constant
     * amount of work is done.
     * 
     * Space Complexity: O(n), the dynamic programming array 'dp' has a size of 'n +
     * 1' to store solutions to subproblems, space complexity is linear with respect
     * to the input size 'n'.
     * 
     * @param n The number of stairs to climb.
     * @return The number of distinct ways to climb the stairs.
     */
    public int climbStairs(int n) {
        // Initialize an array to store solutions to subproblems
        int[] dp = new int[n + 1];

        // Base cases: There is 1 way to climb 0 and 1 step
        dp[0] = 1;
        dp[1] = 1;

        // Dynamic programming approach to fill the array - Bottom-up approach
        // (tabulation)
        for (int i = 2; i < n + 1; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }

        // Return the final result, which represents the number of ways to climb n steps
        return dp[n];
    }
}

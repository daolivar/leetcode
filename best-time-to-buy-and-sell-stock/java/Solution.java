/**
 * Solution that uses the sliding window technique to find the maximum profit
 * when buying and selling stock.
 */
public class Solution {
    /**
     * Calculates the maximum profit that can be obtained by buying and selling
     * stocks.
     * 
     * The sliding window technique is used to keep track of the minimum seen value
     * (representing the lowest stock price) while iterating through the array. The
     * maximum profit is updated by comparing the current profit (difference between
     * the current price and minSeenValue) with the existing maxProfit.
     * 
     * Time Complexity: O(n) where n is the number of elements in the input array
     * prices. The algorithm iterates through the array once, performing
     * constant-time operations at each step.
     * 
     * Space Complexity: O(1) the algorithm uses a constant amount of extra space
     * regardless of the input size.
     * 
     * @param prices An array of stock prices where prices[i] is the price of a
     *               given stock on the ith day.
     * @return The maximum profit that can be obtained by buying and selling stocks.
     */
    public int maxProfit(int[] prices) {
        // Initialize the maximum profit to 0 and the minimum seen value to the maximum
        // possible integer value.
        int maxProfit = 0;
        int minSeenValue = Integer.MAX_VALUE;

        // Iterate through the array to find the maximum profit.
        for (int i = 0; i < prices.length; i++) {
            // If the current price is less than the minimum seen value, update the minimum
            // seen value.
            if (prices[i] < minSeenValue) {
                minSeenValue = prices[i];
            } else {
                // Calculate the profit and update the maximum profit if the current profit is
                // greater.
                maxProfit = Math.max(maxProfit, prices[i] - minSeenValue);
            }
        }

        // Return the final maximum profit.
        return maxProfit;
    }
}

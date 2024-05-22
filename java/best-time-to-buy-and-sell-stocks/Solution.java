/**
 * Solution classes for Leetcode problem 121. Best Time to Buy and Sell Stock
 */
public class Solution {

    /**
     * Finds the maximum profit that can be achieved by buying and selling a stock.
     * 
     * This method uses a sliding window approach where the window's left boundary represents
     * the minimum price encountered so far, and the current price represents the right boundary
     * of the window. The profit is calculated dynamically as the window slides.
     * 
     * @param prices an array of integers where each element represents the price of a stock on a given day
     * @return the maximum profit that can be achieved, or 0 if no profit is possible
     */
    public int maxProfit(int[] prices) {
        int minPrice = Integer.MAX_VALUE; // Initialize the minimum price to the highest possible value
        int maxProfit = 0; // Initialize the maximum profit to 0

        // Iterate through each price in the array
        for (int currentPrice : prices) {
            // Update the minimum price if the current price is lower
            if (currentPrice < minPrice) {
                minPrice = currentPrice;
            } else {
                // Calculate the profit if selling at the current price
                int profit = currentPrice - minPrice;
                // Update the maximum profit if the current profit is higher
                if (profit > maxProfit) {
                    maxProfit = profit;
                }
            }
        }

        return maxProfit; // Return the maximum profit
    }
}
